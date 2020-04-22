/*
 *  Links.js
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
import { Link } from 'react-router-dom';
import { Row, Card, Badge, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faReply, faEdit, faTrash, faCopy } from '@fortawesome/free-solid-svg-icons';
import Pagination from 'react-js-pagination';
import Tags from './Tags';
import { AppContext } from './AppProvider';

export default class Links extends React.Component {
    static contextType = AppContext;
    firstTime = true;

    loadLinks(top = false) {
        if (top) {
            window.location.href = "#links";
        }

        const path = this.props.match.path;
        const slug = this.props.match.params.slug;

        if (path === "/home") {
            this.context.getLinks();
        } else if (path === "/tag/:slug") {
            this.context.getLinksForTag(slug);
        } else if (path === "/search") {
            this.context.getLinksForSearch();
        }
    }

    componentDidMount() {
        if (this.firstTime) {
            this.firstTime = false;
            this.loadLinks();
        }
    }

    componentDidUpdate(prevProps) {
        if (this.props.location !== prevProps.location) {
            this.loadLinks(true);
        }
    }

    handlePageChange(page) {
        const params = new URLSearchParams(this.props.location.search);
        this.context.setSearch('?p='+page+(params.get('q') ? ('&q='+params.get('q')) : ''));
    }

    onDelete(id) {
        this.context.deleteLink(id).then(() => {
            this.loadLinks();
        })
    }

    render() {
        const pag = this.context.state.pagination;
        var title = <h1>Links</h1>;
        if (this.props.match.path === "/tag/:slug") {
            title = (
                <p>
                    <Link to="/home" className="h1">
                        <Badge variant="secondary">
                            {this.props.match.params.slug} ×
                        </Badge>
                    </Link>
                </p>
            );
        } else if (this.props.match.path === "/search") {
            const params = new URLSearchParams(this.props.location.search);
            title = (
                <p>
                    <Link to="/home" className="h1">
                        <Badge variant="primary">
                            🔎 {params.get('q')} ×
                        </Badge>
                    </Link>
                </p>
            )
        }
        return (
            <div id="links">
                {title}
                {this.context.state.links.map((link) => {
                    return <LinkItem link={link} key={link.id} onDelete={this.onDelete.bind(this)} />;
                })}
                <Row className="justify-content-center">
                    <Pagination
                        activePage={pag.currentPage}
                        itemsCountPerPage={pag.perPage}
                        totalItemsCount={pag.perPage * pag.total}
                        pageRangeDisplayed={5}
                        onChange={this.handlePageChange.bind(this)}
                        itemClass="page-item"
                        linkClass="page-link"
                    />
                </Row>
            </div>
        );
    }
}

class LinkItem extends React.Component {
    onDelete() {
        this.props.onDelete(this.props.link.id);
    }

    copyLink() {
        const link = this.props.link.link;
        const el = document.createElement('textarea');
        el.value = link;
        el.setAttribute('readonly', '');
        document.body.appendChild(el);
        el.select();
        document.execCommand('copy');
        document.body.removeChild(el);
    }

    render() {
        const link = this.props.link;
        return (
            <Card className="link">
                <Card.Body>
                        <h5>{link.title}</h5>
                        <Card.Text>
                            <a href={link.link} target="_blank" className="link-url">{link.link}</a>
                        </Card.Text>
                        <Card.Text>
                            <Tags tags={link.tags} className="mb-3" />
                        </Card.Text>
                        <a className="btn btn-primary" href={link.link} target="_blank">
                            <FontAwesomeIcon icon={faReply} /> Go</a>&nbsp;
                        <Button onClick={this.copyLink.bind(this)} variant="dark">
                            <FontAwesomeIcon icon={faCopy} /> Copy</Button>&nbsp;
                        <Link to={"/edit/"+link.id} className="btn btn-warning">
                            <FontAwesomeIcon icon={faEdit} /> Edit</Link>&nbsp;
                        <Link onClick={this.onDelete.bind(this)} className="btn btn-danger">
                            <FontAwesomeIcon icon={faTrash} /> Delete</Link>&nbsp;
                </Card.Body>
            </Card>
        );
    }
}