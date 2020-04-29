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
import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect
} from "react-router-dom";
import { Container, Row, Col, Spinner } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Logo.svg';
import './App.css';

import { AppContext, AppProviderWithRouter } from './AppProvider';
import Header from './Header';
import Footer from './Footer';
import Home from './Home';
import Login from './Login';
import Register from './Register';
import EditLink from './EditLink';
import Site from './Site';
import Landing from './Landing';

export default class App extends React.Component {
    render() {
        return (
			<Router basename={'/app'}>
                <AppProviderWithRouter>
                    <AppContext.Consumer>
                        { context => (
                            <div>
                                <Loading loading={context.state.loading} />
                                <Header />
                                <Container fluid className="py-4">
									<Route exact path="/" component={Landing} />
                                    <Row className="justify-content-center">
                                        <Col lg={8} md={12} sm={12}>
                                            <Switch>
                                                {/*<Route exact path="/">
                                                    {context.state.isAuth ? <Redirect to="/home" /> : <Redirect to="/login" />}
													</Route>*/}
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
                                                <Site />
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
