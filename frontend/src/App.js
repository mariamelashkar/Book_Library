// src/App.js
import React, { useContext, useState, useEffect, useCallback } from 'react';
import { Form, InputGroup, Button, Card, Container, Row, Col, Nav, Tab, Alert, Spinner, Navbar, NavDropdown } from 'react-bootstrap';
import { MyContext, MyProvider } from './components/MyContext';
import AddBookForm from './components/AddBookForm'; // Importing AddBookForm from the correct file
import 'bootstrap/dist/css/bootstrap.min.css';

const Book = ({ book, borrowBook, returnBook }) => (
  <Card className="mb-4">
    <Card.Body>
      <Card.Title>{book.name}</Card.Title>
      <Card.Subtitle className="mb-2 text-muted">Author: {book.author}</Card.Subtitle>
      <Card.Text>Quantity: {book.quantity}</Card.Text>
      <Button variant="primary" onClick={() => borrowBook(book.name)} className="mr-2">Borrow</Button>
      <Button variant="danger" onClick={() => returnBook(book.name)}>Return</Button>
    </Card.Body>
  </Card>
);

const SearchBooks = ({ searchQuery, setSearchQuery }) => (
  <InputGroup className="mb-3">
    <Form.Control
      placeholder="Search by name"
      value={searchQuery}
      onChange={(e) => setSearchQuery(e.target.value)}
    />
  </InputGroup>
);

const BookList = ({ books, borrowBook, returnBook }) => (
  <Row>
    {books.map((book) => (
      <Col md={4} key={book.id}>
        <Book book={book} borrowBook={borrowBook} returnBook={returnBook} />
      </Col>
    ))}
  </Row>
);

const App = () => {
  const { books, setBooks, currentId, setCurrentId } = useContext(MyContext);
  const [searchQuery, setSearchQuery] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchAndRenderData = useCallback(async () => {
    try {
      const response = await fetch('http://localhost:8000/getbook', {
        method: 'GET',
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setBooks(data);
    } catch (error) {
      setError(error);
    } finally {
      setLoading(false);
    }
  }, [setBooks]);

  useEffect(() => {
    fetchAndRenderData();
  }, [fetchAndRenderData]);

  const addBook = async (event) => {
    event.preventDefault();
    const { inlineFormInputName, inlineFormInputAuthor, inlineFormInputQuantity } = event.target.elements;
    const newBook = {
      id: currentId.toString(),
      name: inlineFormInputName.value,
      author: inlineFormInputAuthor.value,
      quantity: parseInt(inlineFormInputQuantity.value, 10),
    };

    try {
      const response = await fetch('http://localhost:8000/createbook', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newBook),
      });

      if (response.status === 201) {
        setCurrentId(currentId + 1);
        fetchAndRenderData();
      } else {
        throw new Error('Error adding book');
      }
    } catch (error) {
      setError(error);
    }
  };

  const borrowBook = async (bookName) => {
    try {
      const response = await fetch(`http://localhost:8000/getbookcheckout/${bookName}`, {
        method: 'PATCH',
      });

      if (response.ok) {
        fetchAndRenderData();
      } else {
        throw new Error('Error borrowing book');
      }
    } catch (error) {
      setError(error);
    }
  };

  const returnBook = async (bookName) => {
    try {
      const response = await fetch(`http://localhost:8000/getbookcheckin/${bookName}`, {
        method: 'PATCH',
      });

      if (response.ok) {
        fetchAndRenderData();
      } else {
        throw new Error('Error returning book');
      }
    } catch (error) {
      setError(error);
    }
  };

  const filteredBooks = (books || []).filter((book) =>
    book.name.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <Container>
      <Navbar bg="light" expand="lg">
        <Navbar.Brand href="#home">Library Management System</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="mr-auto">
            <Nav.Link href="#home">Home</Nav.Link>
            <Nav.Link href="#link">Link</Nav.Link>
            <NavDropdown title="Dropdown" id="basic-nav-dropdown">
              <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
              <NavDropdown.Item href="#action/3.2">Another action</NavDropdown.Item>
              <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
              <NavDropdown.Divider />
              <NavDropdown.Item href="#action/3.4">Separated link</NavDropdown.Item>
            </NavDropdown>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
      <Container fluid className="p-5 bg-light">
        <Row>
          <Col>
            <h1>Welcome to the Library Management System</h1>
            <p>
              This is a simple application to manage your library books. Add, search, and borrow books effortlessly.
            </p>
          </Col>
        </Row>
      </Container>
      <Tab.Container defaultActiveKey="addBook">
        <Row>
          <Col sm={3}>
            <Nav variant="pills" className="flex-column">
              <Nav.Item>
                <Nav.Link eventKey="addBook">Add Book</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link eventKey="searchBooks">Search Books</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link eventKey="bookList">Book List</Nav.Link>
              </Nav.Item>
            </Nav>
          </Col>
          <Col sm={9}>
            <Tab.Content>
              <Tab.Pane eventKey="addBook">
                <AddBookForm addBook={addBook} />
              </Tab.Pane>
              <Tab.Pane eventKey="searchBooks">
                <SearchBooks searchQuery={searchQuery} setSearchQuery={setSearchQuery} />
                {loading ? <Spinner animation="border" /> : <BookList books={filteredBooks} borrowBook={borrowBook} returnBook={returnBook} />}
                {error && <Alert variant="danger">{error.message}</Alert>}
              </Tab.Pane>
              <Tab.Pane eventKey="bookList">
                <Button variant="primary" onClick={fetchAndRenderData} className="mb-4">View All Books</Button>
                {loading ? <Spinner animation="border" /> : <BookList books={filteredBooks} borrowBook={borrowBook} returnBook={returnBook} />}
                {error && <Alert variant="danger">{error.message}</Alert>}
              </Tab.Pane>
            </Tab.Content>
          </Col>
        </Row>
      </Tab.Container>
    </Container>
  );
};

const WrappedApp = () => (
  <MyProvider>
    <App />
  </MyProvider>
);

export default WrappedApp;