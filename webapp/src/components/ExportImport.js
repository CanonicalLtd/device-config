/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

import React, {Component} from 'react';
import {T} from "./Utils";
import {Button, Form, Input, Notification} from "@canonical/react-components";
import api from "./api";

class ExportImport extends Component {
    constructor(props) {
        super(props)
        this.state = {
            data: '',
            message: '',
            messageType: '',
        };
    }

    onChangeFile = (e) => {
        e.preventDefault()

        let reader = new FileReader();
        let file = e.target.files[0];

        reader.onload = (upload) => {
            this.setState({data: upload.target.result.split(',')[1]})
        }

        reader.readAsDataURL(file);
    }

    handleExport = (e) => {
        e.preventDefault()
        window.location.href = api.transferExport()
    }

    handleImport = (e) => {
        e.preventDefault()
        // Call API to use the import file
        api.transferImport(this.state.data).then(response => {
            this.setState({message: T('import-success'), messageType: 'positive'})
        }).catch(e => {
            console.log(e.response)
            this.setState({message: T(e.response.data.code) + ": " + e.response.data.message, messageType:'negative'})
        })
    }

    renderMessage() {
        if (!this.state.message) {
            return
        }
        let status = 'Success:'
        if (this.state.messageType==='negative') {
            status = 'Error:'
        }

        return (
            <Notification type={this.state.messageType} status={status}>
                {this.state.message}
            </Notification>
        )
    }

    render() {
        return (
            <div>
                {this.renderMessage()}
                <h2>{T('export-import')}</h2>
                <p>{T('export-import-desc')}</p>

                <div>
                    <Form>
                        <fieldset>
                            <h3>{T('export')}</h3>
                            <p className="help">{T('export-desc')}</p>
                            <Button onClick={this.handleExport}>{T('export')}</Button>
                        </fieldset>
                    </Form>
                </div>

                <div>
                    <Form>
                        <fieldset>
                            <h3>{T('import')}</h3>
                            <p className="help">{T('import-desc')}</p>
                            <Input onChange={this.onChangeFile} type="file" id="file" placeholder={T('file')} label={T('import-file')} />
                            <Button onClick={this.handleImport}>{T('import')}</Button>
                        </fieldset>
                    </Form>
                </div>
            </div>
        );
    }
}

export default ExportImport;