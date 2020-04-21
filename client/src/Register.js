import React from 'react';
import { Col, Button, Card, Form } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import ValidationErrors from './ValidationErrors';
import { AppContext } from './AppProvider';

export default class Register extends React.Component {
    static contextType = AppContext;

    constructor(props) {
        super(props);
        this.state = {
            name: "",
            email: "",
            password: "",
            confirm: ""
        };
        this.handleInput = this.handleInput.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }

    handleInput(event) {
        const target = event.target;
        const value = target.value;
        const name = target.name;
        this.setState({
            [name]: value
        });
    }

    onSubmit() {
        if (this.state.password === "" || this.state.password === this.state.confirm) {
            this.context.register(this.state);
        } else {
            this.context.addError({"tag": "match"});
        }
    }

    render() {
        return (
            <Card>
                <Card.Header className="h5">Register</Card.Header>
                <Card.Body>
                    <ValidationErrors errors={this.context.errors} />
                    <Form.Group className="row">
                        <Form.Label className="col-md-4 col-form-label text-md-right">Name</Form.Label>
                        <Col md={6}>
                            <Form.Control
                                name="name"
                                value={this.state.name}
                                onChange={this.handleInput}
                                required />
                        </Col>
                    </Form.Group>
                    <Form.Group className="row">
                        <Form.Label className="col-md-4 col-form-label text-md-right">E-Mail Address</Form.Label>
                        <Col md={6}>
                            <Form.Control
                                type="email"
                                name="email"
                                value={this.state.email}
                                onChange={this.handleInput}
                                required />
                        </Col>
                    </Form.Group>
                    <Form.Group className="row">
                        <Form.Label className="col-md-4 col-form-label text-md-right">Password</Form.Label>
                        <Col md={6}>
                            <Form.Control
                                type="password"
                                name="password"
                                value={this.state.password}
                                onChange={this.handleInput}
                                required />
                        </Col>
                    </Form.Group>
                    <Form.Group className="row">
                        <Form.Label className="col-md-4 col-form-label text-md-right">Confirm Password</Form.Label>
                        <Col md={6}>
                            <Form.Control
                                type="password"
                                name="confirm"
                                value={this.state.confirm}
                                onChange={this.handleInput}
                                required />
                        </Col>
                    </Form.Group>
                    <Form.Group className="row mb-0">
                        <Col md={{span: 6, offset: 4}}>
                            <Button 
                                variant="primary"
                                onClick={this.onSubmit}>
                                Register
                            </Button>
                        </Col>
                    </Form.Group>
                </Card.Body>
            </Card>
        );
    }
}