import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import { Container, Row, Col, Spinner } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

import { AppContext, AppProviderWithRouter } from './AppProvider';
import Header from './Header';
import Footer from './Footer';
import Home from './Home';
import Login from './Login';
import Register from './Register';
import EditLink from './EditLink';

export default class App extends React.Component {
    render() {
        return (
            <Router>
                <AppProviderWithRouter>
                    <AppContext.Consumer>
                        { context => (
                            <div>
                                <Loading loading={context.state.loading} />
                                <Header />
                                <Container fluid className="py-4">
                                    <Row className="justify-content-center">
                                        <Col lg={8} md={12} sm={12}>
                                            <Switch>
                                                <Route exact path="/">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Redirect to="/login" />}
                                                </Route>
                                                <Route path="/login">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Login />}
                                                </Route>
                                                <Route path="/register">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Register /> }
                                                </Route>
                                                <Route path={["/home", "/tag/:slug", "/search"]} render={props => (
                                                    !context.state.isAuth ? <Redirect to="/login" /> : <Home location={props.location} match={props.match} />
                                                )} />
                                                <Route path="/edit/:id" render={props => (
                                                    !context.state.isAuth ? <Redirect to="/login" /> : <EditLink edit={true} location={props.location} match={props.match} />
                                                )} />
                                            </Switch>
                                        </Col>
                                    </Row>
                                </Container>
                                <Footer />
                            </div>
                        )}
                    </AppContext.Consumer>
                </AppProviderWithRouter>
            </Router>
        );
    }
}

function Loading({...props}) {
    if (props.loading) {
        return (
            <div className="loading">
                <Spinner className="loading-spinner" variant="primary" animation="border" role="status">
                    <span className="sr-only">Loading...</span>
                </Spinner>
            </div>
        );
    } else {
        return <div></div>;
    }
}