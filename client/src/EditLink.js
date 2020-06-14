/*
 *  EditLink.js
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
import { faEdit, faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { Button, Card, Form } from 'react-bootstrap';
import TagsInput from 'react-tagsinput';
import { AppContext } from './AppProvider';
import ValidationErrors from './ValidationErrors';

export default class EditLink extends React.Component {
    static contextType = AppContext;
    constructor(props) {
        super(props)
        this.state = {
            id: null,
            title: "",
            link: "",
            tags: []
        };
        this.handleTags = this.handleTags.bind(this);
        this.handleInput = this.handleInput.bind(this);
        this.onSubmit = this.onSubmit.bind(this);
    }

    componentDidMount() {
        if (this.props.edit) {
            const id = this.props.match.params.id;
            this.context.getLink(id).then(link => {
                if (link) {
                    const tags = link.tags.map(tag => {
                        return tag.name;
                    });
                    this.setState({
                        id: link.id,
                        title: link.title,
                        link: link.link,
                        tags: tags
                    });
                }
            })
        }
    }

    handleTags(tags) {
        this.setState({ tags: tags })
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
        const tags = this.state.tags.join(',');
        if (!this.props.edit) {
            this.context.addLink({
                link: this.state.link,
                tags: tags
            }).then(response => {
                if (response) {
                    this.setState({
                        title: "",
                        link: "",
                        tags: []
                    });
                    this.context.goHome();
                }
            });
        } else {
            this.context.updateLink(this.state.id, {
                title: this.state.title,
                link: this.state.link,
                tags: tags
            }).then(() => {
                this.context.goBack();
            })
        }
    }

    render() {
        return (
            <div className={this.props.edit && "container"}>
                {this.props.edit &&
                    <p>
                        <Button variant="warning"
                            size="lg"
                            onClick={this.context.goBack}>Go Back</Button>
                    </p>
                }
                <Card>
                    <Card.Header className="h5">Add Link</Card.Header>
                    <Card.Body>
                        <ValidationErrors />
                        {this.props.edit &&
                            <Form.Group>
                                <Form.Label>Title</Form.Label>
                                <Form.Control
                                    name="title"
                                    placeholder="Example Website"
                                    value={this.state.title}
                                    onChange={this.handleInput}
                                />
                            </Form.Group>
                        }
                        <Form.Group>
                            <Form.Label>Link</Form.Label>
                            <Form.Control
                                type="url"
                                name="link"
                                placeholder="https://example.com/"
                                value={this.state.link}
                                onChange={this.handleInput}
                            />
                        </Form.Group>
                        <Form.Group>
                            <Form.Label>Tags</Form.Label>
                            <TagsInput
                                value={this.state.tags}
                                onChange={this.handleTags}
                            />
                        </Form.Group>
                        {this.props.edit ?
                            <Button variant="warning" onClick={this.onSubmit}>
                                <FontAwesomeIcon icon={faEdit} /> Edit link</Button> :
                            <Button variant="primary" onClick={this.onSubmit}>
                                <FontAwesomeIcon icon={faPlus} /> Add link</Button>
                        }
                    </Card.Body>
                </Card>
            </div>
        );
    }
}
