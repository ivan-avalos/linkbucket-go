import React from 'react';
import { Card, Form, Button } from 'react-bootstrap';
import TagsInput from 'react-tagsinput';
import ValidationErrors from './ValidationErrors';
import { AppContext } from './AppProvider';

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
                    const tags = link.tags.map(tag =>{
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
        this.setState({tags: tags})
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
            <div>
                { this.props.edit && 
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
                        { this.props.edit &&
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
                        { this.props.edit ?
                             <Button variant="warning" onClick={this.onSubmit}>Edit Link</Button> :
                             <Button variant="primary" onClick={this.onSubmit}>Add Link</Button>
                        }
                    </Card.Body>
                </Card>
            </div>
        );
    }
}