package handler

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/dto"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/service"
	mock_service "github.com/ivanovamir/Pure-architecture-REST-API/internal/service/mocks"
	"github.com/ivanovamir/Pure-architecture-REST-API/internal/transport/handler"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_GetAllBooks(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBook)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK empty result",
			mockBehavior: func(s *mock_service.MockBook) {
				s.EXPECT().GetAllBooks(context.Background()).Return([]*dto.Book{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":[]}`,
		},
		{
			name: "ERROR scanning rows",
			mockBehavior: func(s *mock_service.MockBook) {
				s.EXPECT().GetAllBooks(context.Background()).Return(nil, fmt.Errorf("error occurred scanning rows"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"error occurred scanning rows"}`,
		},
		{
			name: "ERROR parsing rows",
			mockBehavior: func(s *mock_service.MockBook) {
				s.EXPECT().GetAllBooks(context.Background()).Return([]*dto.Book{}, fmt.Errorf("error occurred parsing rows"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"error occurred parsing rows"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			books := mock_service.NewMockBook(c)

			testCase.mockBehavior(books)

			service := &service.Service{
				Book: books,
			}

			router := httprouter.New()

			handler := handler.NewHttpHandler(router, service)
			router.GET("/books", handler.GetAllBooks)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/books", nil)
			// Perform Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())

		})
	}
}

func TestHandler_GetBookByID(t *testing.T) {
	type mockBehavior func(s *mock_service.MockBook, bookId int)

	testTable := []struct {
		name                 string
		inputParams          int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK good result #1",
			inputParams: 1,
			mockBehavior: func(s *mock_service.MockBook, bookId int) {
				s.EXPECT().GetBookByID(context.Background(), bookId).Return(&dto.Book{
					Id:    "1",
					Title: "Евгений Онегин",
					Year:  "1860",
					Author: dto.Author{
						Id:   "1",
						Name: "Пушкин А.C.",
					},
					Genre: dto.Genre{
						Id:    "2",
						Title: "Поэма",
					},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":{"id":"1","title":"Евгений Онегин","year":"1860","author":{"id":"1","name":"Пушкин А.C."},"genre":{"id":"2","title":"Поэма"}}}`,
		},
		{
			name:        "OK good result #2",
			inputParams: 2,
			mockBehavior: func(s *mock_service.MockBook, bookId int) {
				s.EXPECT().GetBookByID(context.Background(), bookId).Return(&dto.Book{
					Id:    "2",
					Title: "Руслан Людмила",
					Year:  "1875",
					Author: dto.Author{
						Id:   "1",
						Name: "Пушкин А.C.",
					},
					Genre: dto.Genre{
						Id:    "1",
						Title: "Роман",
					},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":{"id":"2","title":"Руслан Людмила","year":"1875","author":{"id":"1","name":"Пушкин А.C."},"genre":{"id":"1","title":"Роман"}}}`,
		},
		{
			name:        "OK good result #3",
			inputParams: 3,
			mockBehavior: func(s *mock_service.MockBook, bookId int) {
				s.EXPECT().GetBookByID(context.Background(), bookId).Return(&dto.Book{
					Id:    "3",
					Title: "Преступление и наказание",
					Year:  "1855",
					Author: dto.Author{
						Id:   "3",
						Name: "Достоевский Ф.М.",
					},
					Genre: dto.Genre{
						Id:    "2",
						Title: "Поэма",
					},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":{"id":"3","title":"Преступление и наказание","year":"1855","author":{"id":"3","name":"Достоевский Ф.М."},"genre":{"id":"2","title":"Поэма"}}}`,
		},
		{
			name:        "OK empty result",
			inputParams: 0,
			mockBehavior: func(s *mock_service.MockBook, bookId int) {
				s.EXPECT().GetBookByID(context.Background(), bookId).Return(&dto.Book{
					Id:     "",
					Title:  "",
					Year:   "",
					Author: dto.Author{},
					Genre:  dto.Genre{},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":{}}`,
		},
		{
			name:        "ERROR scanning rows",
			inputParams: 0,
			mockBehavior: func(s *mock_service.MockBook, bookId int) {
				s.EXPECT().GetBookByID(context.Background(), bookId).Return(nil, fmt.Errorf("%s", "error occurred scanning rows"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"error occurred scanning rows"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			bookById := mock_service.NewMockBook(c)

			testCase.mockBehavior(bookById, testCase.inputParams)

			services := &service.Service{
				Book: bookById,
			}
			// Create http_router
			router := httprouter.New()
			handler := handler.NewHttpHandler(router, services)

			router.GET("/book", handler.GetBookByID)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/book", nil)
			query := req.URL.Query()
			query.Add("book_id", strconv.Itoa(testCase.inputParams))
			req.URL.RawQuery = query.Encode()

			// Perform Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
