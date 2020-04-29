import React from 'react';
import { Link } from 'react-router-dom';
import Logo from './Logo.svg';
import { Card, Container, Button } from 'react-bootstrap';

export default class Landing extends React.Component {
	render() {
		return (
			<div className="text-center landing-block-1">
				<Container>
					<h1 className="display-4 landing-title-1">
						The free and open source bookmark manager.
					</h1>
					<br/>
					<h2 className="landing-title-2">It is supposed to respect your privacy.</h2>
					<br/>
					<Link to="/register" className="btn btn-lg btn-primary">Join</Link>&nbsp;
					<Link to="/login" className="btn btn-lg btn-link">Login</Link>
				</Container>
			</div>
		)
	}
}
