/*
 *  Register.js
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
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