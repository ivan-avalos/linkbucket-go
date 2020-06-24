/*
 *  App.js
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
import 'bootstrap/dist/css/bootstrap.min.css';
import React from 'react';
import { Col, Container, Row, Spinner } from 'react-bootstrap';
import {
    BrowserRouter as Router,


    Redirect, Route, Switch
} from "react-router-dom";
import './App.css';
import { AppContext, AppProviderWithRouter } from './AppProvider';
import EditLink from './EditLink';
import Footer from './Footer';
import Header from './Header';
import Home from './Home';
import Landing from './Landing';
import Login from './Login';
import './Logo.svg';
import Register from './Register';
import Site from './Site';


export default class App extends React.Component {
    render() {
        return (
            <Router basename={'/app'}>
                <AppProviderWithRouter>
                    <AppContext.Consumer>
                        {context => (
                            <div>
                                <Loading loading={context.state.loading} />
                                <Header />
                                <Container fluid className="py-4 px-0 mx-0">
                                    <Route exact path="/" component={Landing} />
                                    <Row style={{ marginLeft: 0, marginRight: 0 }} className="justify-content-center">
                                        <Col lg={8} md={12} sm={12} style={{ paddingLeft: 0, paddingRight: 0 }}>
                                            <Switch>
                                                <Route exact path="/">
                                                    {context.state.isAuth && <Redirect to="/home" />}
                                                </Route>
                                                <Route path="/login">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Login />}
                                                </Route>
                                                <Route path="/register">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Register />}
                                                </Route>
                                                <Route path={["/home", "/tag/:slug", "/search"]} render={props => (
                                                    !context.state.isAuth ? <Redirect to="/login" /> : <Home location={props.location} match={props.match} />
                                                )} />
                                                <Route path="/edit/:id" render={props => (
                                                    !context.state.isAuth ? <Redirect to="/login" /> : <EditLink edit={true} location={props.location} match={props.match} />
                                                )} />
                                                <Route path="/site">
                                                    <Site />
                                                </Route>
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

function Loading({ ...props }) {
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
