/*
 *  Home.js
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
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { Card, Container, Form, InputGroup } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { AppContext } from './AppProvider';
import EditLink from './EditLink';
import Links from './Links';
import Tags from './Tags';

export default class Home extends React.Component {
    static contextType = AppContext;

    componentDidMount() {
        this.context.getTags();
    }

    componentDidUpdate(prevProps) {
        if (prevProps.location !== this.props.location) {
            this.context.getTags();
        }
    }

    render() {
        return (
            <Container>
                <EditLink edit={false} />
                <br />
                <Search />
                <br />
                <Links location={this.props.location} match={this.props.match} />
            </Container>
        );
    }
}

class Search extends React.Component {
    static contextType = AppContext;
    state = { q: "" }

    handleInput(event) {
        const target = event.target;
        const value = target.value;
        this.setState({
            q: value
        });
    }

    render() {
        return (
            <Card>
                <Card.Header className="h5">Search</Card.Header>
                <Card.Body>
                    <p>
                        <InputGroup>
                            <Form.Control
                                placeholder="Enter search query"
                                aria-label="Enter search query"
                                onChange={this.handleInput.bind(this)}
                            />
                            <InputGroup.Append>
                                <Link class="btn btn-primary" type="submit" to={"/search?q=" + this.state.q}>
                                    <FontAwesomeIcon icon={faSearch} /> Search</Link>
                            </InputGroup.Append>
                        </InputGroup>
                    </p>
                    <h4>Tags</h4>
                    <Tags count tags={this.context.state.tags} />
                </Card.Body>
            </Card>
        );
    }
}
