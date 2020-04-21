import React from 'react';
import { Link } from 'react-router-dom';
import { Container, Card, Form, Button, InputGroup } from 'react-bootstrap';
import EditLink from './EditLink';
import Tags from './Tags';
import Links from './Links';
import { AppContext } from './AppProvider';

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
    state = {q: ""}

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
                                {/*<Button variant="primary" type="submit">Search</Button>*/}
                                <Link class="btn btn-primary" type="submit" to={"/search?q="+this.state.q}>Search</Link>
                            </InputGroup.Append>
                        </InputGroup>
                    </p>
                    <h4>Tags</h4>
                    <Tags tags={this.context.state.tags} />
                </Card.Body>
            </Card>
        );
    }
}