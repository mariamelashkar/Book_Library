// src/AddBookForm.js
import React from 'react';
import { Form, Col, Button } from 'react-bootstrap';

const AddBookForm = ({ addBook }) => (
  <Form onSubmit={addBook}>
    <Form.Row className="align-items-center">
      <Col sm={3} className="my-1">
        <Form.Label htmlFor="inlineFormInputName" srOnly>
          Book Name
        </Form.Label>
        <Form.Control id="inlineFormInputName" placeholder="Book Name" required />
      </Col>
      <Col sm={3} className="my-1">
        <Form.Label htmlFor="inlineFormInputAuthor" srOnly>
          Author
        </Form.Label>
        <Form.Control id="inlineFormInputAuthor" placeholder="Author" required />
      </Col>
      <Col sm={2} className="my-1">
        <Form.Label htmlFor="inlineFormInputQuantity" srOnly>
          Quantity
        </Form.Label>
        <Form.Control id="inlineFormInputQuantity" type="number" placeholder="Quantity" required />
      </Col>
      <Col xs="auto" className="my-1">
        <Button type="submit">Add Book</Button>
      </Col>
    </Form.Row>
  </Form>
);

export default AddBookForm;
