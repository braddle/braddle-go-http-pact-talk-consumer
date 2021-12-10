package book_test

import (
	book "book-client"
	"errors"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ClientSuite struct {
	suite.Suite
}

var pact *dsl.Pact

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupSuite() {
	pact = &dsl.Pact{
		Consumer: "MarksBookClient",
		Provider: "BookService",
		PactDir:  "./pacts",
	}
}

func (s *ClientSuite) TearDownSuite() {
	pact.Teardown()
}

func (s *ClientSuite) TestRetrievingBookThatDoesNotExist() {
	pact.AddInteraction().
		Given("There is not a book with ISBN 123456789").
		UponReceiving("A GET request for book with ISBN 123456789").
		WithRequest(
			dsl.Request{
				Method: http.MethodGet,
				Path:   dsl.String("/book/123-4567890123"),
				Headers: dsl.MapMatcher{
					"Accept": dsl.String("application/json"),
				},
			},
		).
		WillRespondWith(
			dsl.Response{
				Status: http.StatusNotFound,
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.String("application/json"),
				},
				Body: dsl.Match(book.JsonError{}),
			},
		)

	test := func() error {
		c := book.NewClient(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
		_, err := c.GetBook("123-4567890123")

		s.Error(err)
		s.Equal(err.Error(), "No book with ISBN 123-4567890123")

		if s.T().Failed() {
			return errors.New("failed validating error response")
		}

		return nil
	}

	s.NoError(pact.Verify(test))
}

func (s *ClientSuite) TestRetrievingBookThatDoesExist() {
	s.False(s.T().Failed())
}

func (s *ClientSuite) TestCreatingBook() {

}