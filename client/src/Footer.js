import React from 'react';
import { Container } from 'react-bootstrap';

export default class Footer extends React.Component {
    render() {
        return (
            <footer className="footer py-3">
                <Container fluid>
                    <span className="text-muted">LINKBUCKET</span>
                </Container>
            </footer>
        );
    }
}