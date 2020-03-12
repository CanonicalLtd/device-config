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
import api from "./api";
import {formatError, T} from "./Utils";
import AlertBox from "./AlertBox";

class Services extends Component {
    constructor(props) {
        super(props)
        this.state = {
            services: [],
            error: '',
        };
    }

    componentDidMount() {
        this.getServices()
    }

    getServices = () => {
        api.servicesGet().then(response => {
            console.log('---', response)
            this.setState({services: response.data.services, error: ''})
        })
        .catch(e => {
            console.log('---', e.response)
            this.setState({error: formatError(e.response.data)});
        })
    }

    renderError() {
        if (this.state.error) {
            return <AlertBox message={this.state.error}/>
        }
    }

    render() {
        return (
            <div>
                <h2>{T('service-status')}</h2>

                <section className="row">
                    {this.renderError()}
                    <table>
                        <thead>
                            <tr>
                                <th>{T('name')}</th>
                                <th>{T('service')}</th>
                                <th className="xsmall u-align-text--center">{T('enabled')}</th>
                                <th className="xsmall u-align-text--center">{T('active')}</th>
                            </tr>
                        </thead>
                        <tbody>
                        {this.state.services.map(srv => {
                            return (<tr>
                                <td>{srv.snap}</td>
                                <td>{srv.name}</td>
                                <td>{srv.enabled ? <div className="led-green"></div> : <div className="led-yellow"></div> }</td>
                                <td>{srv.active ? <div className="led-green"></div> : <div className="led-red"></div>}</td>
                            </tr>)
                        })}
                        </tbody>
                    </table>
                </section>
            </div>
        );
    }
}

export default Services;