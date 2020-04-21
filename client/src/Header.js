/*
 *  Header.js
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
import { Link } from "react-router-dom";
import { Navbar, Nav, NavDropdown } from 'react-bootstrap';
import { AppContext } from './AppProvider';
import Logo from './Logo.svg';

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
                <Navbar.Brand className="mr-auto">
                    <img src={Logo}
                        height="35"
                        className="d-inline-block align-top"
                        alt="Linkbucket" />
                </Navbar.Brand>
                {links}
            </Navbar>
        );
    }
}