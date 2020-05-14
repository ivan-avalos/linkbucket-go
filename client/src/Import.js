import React from 'react';
import { Button, Modal, Form, FormGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBan, faDownload } from '@fortawesome/free-solid-svg-icons';
import { AppContext } from './AppProvider';

export class Import extends React.Component {
    static contextType = AppContext;
    file = null;

    setFile(event) {
        const target = event.target;
        this.file = target.files[0];
    }

    onSubmit() {
        console.log("onSubmit()");
        if (this.file !== null) {
            console.log(this.file);
            this.context.importOld(this.file).then(success => {
                if (success) {
                    this.props.handleClose();
                }
            });
        }
    }

    render() {
        return (
            <Modal show={this.props.show} onHide={this.props.handleClose}>
                <Modal.Header>
                    <Modal.Title>Import Bookmarks</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <p>
                        Importing bookmarks is an asynchronous task, that means that, after you are redirected to home, you will have to refresh in a while in order to see all the bookmarks. That "while" will depend in how many bookmarks you are importing.
                    </p>
                    <p>
                        <b>Duplicate links won't be inserted.</b>
                    </p>
                    <FormGroup>
                        <Form.Check type="radio" label="Import from Pocket." disabled />
                        <Form.Check type="radio" label="Import from Linkbucket (Laravel)." checked />
                    </FormGroup>
                    <FormGroup>
                        <Form.File id="import" label="Select JSON file" onChange={this.setFile.bind(this)} />
                    </FormGroup>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="primary" onClick={this.props.handleClose.bind(this)}>
                        <FontAwesomeIcon icon={faBan} /> Cancel</Button>
                    <Button variant="danger" onClick={this.onSubmit.bind(this)}>
                        <FontAwesomeIcon icon={faDownload} /> Import</Button>
                </Modal.Footer>
            </Modal>
        );
    }
}
