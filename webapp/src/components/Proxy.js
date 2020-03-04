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
import {formatError, T} from "./Utils";
import AlertBox from "./AlertBox";
import api from "./api";

class Proxy extends Component {
    constructor(props) {
        super(props)
        this.state = {
            proxy: {},
            error: '',
            message: '',
        };
    }

    componentDidMount() {
        this.getProxyConfig()
    }

    getProxyConfig = () => {
        api.proxyGet().then(response => {
            this.setState({proxy: response.data.proxy, message: ''})
        })
        .catch(e => {
            this.setState({message: formatError(e.response.data)});
        })
    }

    handleHTTPChange = (e) => {
        e.preventDefault()
        let iface = this.state.proxy
        iface.http = e.target.value
        this.setState({proxy: iface})
    }

    handleHTTPSChange = (e) => {
        e.preventDefault()
        let iface = this.state.proxy
        iface.https = e.target.value
        this.setState({proxy: iface})
    }

    handleFTPChange = (e) => {
        e.preventDefault()
        let iface = this.state.proxy
        iface.ftp = e.target.value
        this.setState({proxy: iface})
    }

    handleSave = (e) => {
        e.preventDefault()

        // Save the proxy config
        api.proxyUpdate(this.state.proxy).then(response => {
            this.setState({message: T('proxy-updated'), error: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }


    renderMessage() {
        if (this.state.message) {
            return <AlertBox type="positive" message={this.state.message}/>
        }
    }

    renderError() {
        if (this.state.error) {
            return <AlertBox message={this.state.error}/>
        }
    }
    render() {
        return (
            <div>
                <h2>{T('proxy-config')}</h2>

                <section className="row">
                    {this.renderMessage()}
                    {this.renderError()}
                    <form>
                        <fieldset>
                        <label htmlFor={"http"}>{T('http')}:</label>
                        <input name="http" type="text" value={this.state.proxy.http} onChange={this.handleHTTPChange}
                            placeholder={T('proxy-help')}/>
                        <label htmlFor={"https"}>{T('https')}:</label>
                        <input name="https" type="text" value={this.state.proxy.https} onChange={this.handleHTTPSChange}
                               placeholder={T('proxy-help')}/>
                        <label htmlFor={"ftp"}>{T('ftp')}:</label>
                        <input name="ftp" type="text" value={this.state.proxy.ftp} onChange={this.handleFTPChange}
                               placeholder={T('proxy-help')}/>

                        <button onClick={this.handleSave} className="p-button--positive">{T('save')}</button>
                        </fieldset>
                    </form>
                </section>

            </div>
        );
    }
}

export default Proxy;