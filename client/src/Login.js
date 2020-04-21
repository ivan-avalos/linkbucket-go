import React from 'react';
import { Card, Form, Col, Button } from 'react-bootstrap';
import { AppContext } from './AppProvider';

import ValidationErrors from './ValidationErrors';

export default class Login extends React.Component {
    static contextType = AppContext;

    constructor(props) {
        super(props);
        this.state = {
            email: "",
            password: ""
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
        this.context.login(this.state);
    }

    render() {
        return (
            <Card>
                <Card.Header className="h5">Login</Card.Header>
                <Card.Body>
                    <ValidationErrors />
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
                    <Form.Group className="row mb-0">
                        <Col md={{span: 6, offset: 4}}>
                            <Button 
                                variant="primary"
                                onClick={this.onSubmit}>
                                Login
                            </Button>
                        </Col>
                    </Form.Group>
                </Card.Body>
            </Card>
        );
    }
}