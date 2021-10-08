package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

// для парсинга времени в нужном формате
const (
	layout1 = `"` + "2006-01-02" + `"`
	layout2 = "2006-01-02"
	layout3 = "2006-01"
)

func main() {
	service := NewService()
	handler := NewHandler(service)

	mux := handler.InitRoutes()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		ErrorResponse(nil, err.Error(), http.StatusInternalServerError)
	}
}

type Event struct {
	UserID      int    `json:"user_id"`
	EventID     int    `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        Date   `json:"date"`
}

// Date - свой формат даты
type Date struct {
	time.Time
}

// Events - хранилище все событий
var Events []Event

// UnmarshalJSON - функция для возможности ввода даты в нужном нам формате
func (t *Date) UnmarshalJSON(date []byte) error {
	if string(date) == "" || string(date) == `""` || string(date) == "null" {
		*t = Date{time.Now()}
		return nil
	}
	tm, err := time.Parse(layout1, string(date))
	*t = Date{tm}
	return err
}

// IEvent - содержит поддерживаемые методы
type IEvent interface {
	CreateEvent(userID int, eventID int, title string, description string, date Date) error
	UpdateEvent(userID int, eventID int, title string, description string, date Date) error
	DeleteEvent(eventID int) error
	EventsForDay(date Date, userID int) ([]Event, error)
	EventsForWeek(date Date, userID int) ([]Event, error)
	EventsForMonth(date Date, userID int) ([]Event, error)
}

// Service - сервис, отвечающий за выполнение методов
type Service struct {
	IEvent
}

func NewService() *Service {
	return &Service{NewEventsManager()}
}

// EventsManager - менеджер, осуществляющий методы
type EventsManager struct {
}

func NewEventsManager() *EventsManager {
	return &EventsManager{}
}

// CreateEvent - функция непосредствееного создания события менеджером
func (e *EventsManager) CreateEvent(userID int, eventID int, title string, description string, date Date) error {
	event := Event{UserID: userID, EventID: eventID, Title: title, Description: description, Date: date}
	Events = append(Events, event)
	return nil
}

// UpdateEvent - функция изменения информации о событии менеджером
func (e *EventsManager) UpdateEvent(userID int, eventID int, title string, description string, date Date) error {
	// проверка на существование события
	var index = -1
	for i, e := range Events {
		if e.EventID == eventID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("event doesn't exist")
	}

	Events[index].UserID = userID
	Events[index].Date = date
	Events[index].Title = title
	Events[index].Description = description
	return nil
}

// DeleteEvent - функция удаление события менеджером
func (e *EventsManager) DeleteEvent(eventID int) error {
	// проверка на существование события
	var index int
	for i, e := range Events {
		if e.EventID == eventID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("event doesn't exist")
	}
	Events[index] = Events[len(Events)-1]
	Events = Events[:len(Events)-1]
	return nil
}

// EventsForDay - функция вывода событий за день менеджером
func (e *EventsManager) EventsForDay(date Date, userID int) ([]Event, error) {
	var DayEvents []Event
	for _, e := range Events {
		if e.Date == date && e.UserID == userID {
			DayEvents = append(DayEvents, e)
		}
	}
	return DayEvents, nil
}

// EventsForWeek - функция вывода событий за неделю менеджером
func (e *EventsManager) EventsForWeek(date Date, userID int) ([]Event, error) {
	var WeekEvents []Event
	for _, e := range Events {
		year1, week1 := date.ISOWeek()
		year2, week2 := e.Date.ISOWeek()
		if year1 == year2 && week1 == week2 && e.UserID == userID {
			WeekEvents = append(WeekEvents, e)
		}
	}
	return WeekEvents, nil
}

// EventsForMonth - функция вывода событий за месяц менеджером
func (e *EventsManager) EventsForMonth(date Date, userID int) ([]Event, error) {
	var MonthEvents []Event
	for _, e := range Events {
		if date.Month() == e.Date.Month() && date.Year() == e.Date.Year() && e.UserID == userID {
			MonthEvents = append(MonthEvents, e)
		}
	}
	return MonthEvents, nil
}

// errorResponse - структура вывода ошибок
type errorResponse struct {
	Err string `json:"error"`
}

// ErrorResponse - функция вывода ошибок
func ErrorResponse(w http.ResponseWriter, err string, status int) {
	jsonError, _ := json.Marshal(&errorResponse{err})
	http.Error(w, string(jsonError), status)
}

// resultResponse - структура вывода результатов
type resultResponse struct {
	Result []Event `json:"result"`
}

// ResultResponse - функция вывода результатов
func ResultResponse(w http.ResponseWriter, res []Event) {
	jsonError, _ := json.Marshal(&resultResponse{res})
	_, err := w.Write(jsonError)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadGateway)
	}
}

// Handler - обработчик сервисов
type Handler struct {
	Services *Service
}

func NewHandler(services *Service) *Handler {
	return &Handler{Services: services}
}

// jsonDecode - функция обработки файла json для представления события
func (h *Handler) jsonDecode(w http.ResponseWriter, r *http.Request) *Event {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return &event
}

// createEvent(Handler) - функция создания события обработчиком и конечный вывод
func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	event := h.jsonDecode(w, r)
	if event == nil {
		return
	}
	err := h.Services.CreateEvent(event.UserID, event.EventID, event.Title, event.Description, event.Date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, Events)
}

// updateEvent(Handler) - функция обновления события обработчиком и конечный вывод
func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	event := h.jsonDecode(w, r)
	if event == nil {
		return
	}
	err := h.Services.UpdateEvent(event.UserID, event.EventID, event.Title, event.Description, event.Date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, Events)
}

// deleteEvent(Handler) - функция удаления события обработчиком и конечный вывод
func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	event := h.jsonDecode(w, r)
	if event == nil {
		return
	}
	err := h.Services.DeleteEvent(event.EventID)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, Events)
}

// eventsForDay(Handler) - функция вывода событий за день обработчиком
func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDstr)
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(layout2, date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	events, err := h.Services.EventsForDay(Date{eventTime}, userID)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, events)
}

// eventsForWeek(Handler) - функция вывода событий за неделю обработчиком
func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDstr)
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(layout2, date)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	events, err := h.Services.EventsForWeek(Date{eventTime}, userID)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, events)
}

// eventsForMonth(Handler) - функция вывода событий за месяц обработчиком
func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDstr)
	date := r.URL.Query().Get("date")
	eventTime, err := time.Parse(layout3, date)
	if err != nil {
		eventTime, err = time.Parse(layout2, date)
		if err != nil {
			ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	events, err := h.Services.EventsForMonth(Date{eventTime}, userID)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	ResultResponse(w, events)
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := &http.ServeMux{}

	mux.HandleFunc("/create_event", LogMiddleware(http.HandlerFunc(h.createEvent)))
	mux.HandleFunc("/update_event", LogMiddleware(http.HandlerFunc(h.updateEvent)))
	mux.HandleFunc("/delete_event", LogMiddleware(http.HandlerFunc(h.deleteEvent)))
	mux.HandleFunc("/events_for_day", LogMiddleware(http.HandlerFunc(h.eventsForDay)))
	mux.HandleFunc("/events_for_week", LogMiddleware(http.HandlerFunc(h.eventsForWeek)))
	mux.HandleFunc("/events_for_month", LogMiddleware(http.HandlerFunc(h.eventsForMonth)))

	return mux
}

// LogMiddleware - функция логирования запросов
func LogMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: Method " + r.Method + ", Url " + r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
