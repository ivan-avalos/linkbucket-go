import React from 'react';
import { Link } from "react-router-dom";
import { Navbar, Nav, NavDropdown } from 'react-bootstrap';
import { AppContext } from './AppProvider';

export default class Header extends React.Component {
    static contextType = AppContext;
    render() {
        let links;
        if (!this.context.state.isAuth) {
            links = <Nav>
                <Link to="/login" className="nav-link">Login</Link>
                <Link to="/register" className="nav-link">Register</Link>
            </Nav>;
        } else {
            links = <NavDropdown alignRight title={this.context.state.user.name+" "}>
                <NavDropdown.Item onClick={this.context.import}>
                    Import
                </NavDropdown.Item>
                <NavDropdown.Item onClick={this.context.import}>
                    Export
                </NavDropdown.Item>
                <NavDropdown.Item onClick={this.context.logout}>
                    Logout
                </NavDropdown.Item>
            </NavDropdown>
        }
        return (
            <Navbar bg="light" variant="light">
                <Navbar.Brand className="mr-auto">Linkbucket</Navbar.Brand>
                {links}
            </Navbar>
        );
    }
}