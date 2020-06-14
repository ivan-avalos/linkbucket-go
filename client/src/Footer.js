/*
 *  Footer.js
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
import { Col, Container, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';

export default class Footer extends React.Component {
    render() {
        return (
            <footer>
                <Container className="py-5">
                    <Row>
                        <Col md={6}>
                            <p class="mb-0 text-uppercase font-weight-bold small">
                                <Link to="/site/about" class="text-primary pr-3">About</Link>
                                <Link to="/site/open-source" class="text-primary pr-3">FOSS</Link>
                                <Link to="/site/terms" class="text-primary pr-3"> Terms </Link>
                                <Link to="/site/privacy" class="text-primary pr-3"> Privacy </Link>
                                <Link href="/site/api" class="text-primary pr-3">API</Link>
                            </p>
                        </Col>
                        <Col md={6}>
                            <p class="mb-0 text-uppercase font-weight-bold small">
                                <a href="https://github.com/ivan-avalos/linkbucket-go" class="text-muted" rel="noopener">Powered by Linkbucket</a>
                            </p>
                        </Col>
                    </Row>
                </Container>
            </footer>
        );
    }
}
