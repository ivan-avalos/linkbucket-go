/*
 *  Login.js
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
import { Route, Link } from 'react-router-dom';
import { Card, Form, Col, Button } from 'react-bootstrap';
import { AppContext } from './AppProvider';

export default class Site extends React.Component {
    render() {
        return (
			<div>
				<Route path="/site/about">
					<Card>
						<Card.Header className="h5">
							About
						</Card.Header>
						<Card.Body className="lead">
							Linkbucket is a free and open-source online bookmark manager focused on simplicity and minimalism.
						</Card.Body>
					</Card>
				</Route>
				<Route path="/site/open-source">
					<Card>
						<Card.Header className="h5">
							Open Source
						</Card.Header>
						<Card.Body className="lead">
							The software that powers this website is called Linkbucket and anyone can download the&nbsp;
							<a className="lead" href="https://github.com/ivan-avalos/linkbucket-go">source code</a> and run their own instance! 
							It is written in Go and React and it is totally self-hostable, so you can run your own instance.
						</Card.Body>
					</Card>
				</Route>
			</div>
        );
    }
}
