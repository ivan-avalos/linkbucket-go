import React from 'react';
import { Link } from 'react-router-dom';
import { Row, Card, Badge } from 'react-bootstrap';
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
                    return <LinkItem link={link} key={link.id} />;
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
    render() {
        const link = this.props.link;
        return (
            <Card className="link">
                <Card.Body>
                        <h5>{link.title}</h5>
                        <Card.Text>
                            <a href={link.link} className="link-url">{link.link}</a>
                        </Card.Text>
                        <Card.Text>
                            <Tags tags={link.tags} className="mb-3" />
                        </Card.Text>
                        <a className="btn btn-primary" href={link.link}>Go</a>&nbsp;
                        <Link to={"/edit/"+link.id} className="btn btn-warning">Edit</Link>&nbsp;
                        <Link to={"/delete/"+link.id} className="btn btn-danger">Delete</Link>
                </Card.Body>
            </Card>
        );
    }
}