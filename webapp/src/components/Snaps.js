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
import api from "./api";
import AlertBox from "./AlertBox";

class Snaps extends Component {
    constructor(props) {
        super(props)
        this.state = {
            snaps: [],
            settings: {},
        };

    }

    componentDidMount() {
        this.getSnaps()
    }

    getSnaps = () => {
        api.snapsGet().then(response => {
            this.setState({snaps: response.data.records, error: '', message: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }

    findSnap(name) {
        return this.state.snaps.filter( s => s.name===name)[0]
    }

    handleShowSettings = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');
        if (this.state.snapSettings === snap) {
            this.setState({snapSettings: null, settings: {}})
        } else {
            var s = this.findSnap(snap)
            if (s.config.length===0) {
                s.config = '{}'
            }
            var settings = JSON.stringify( JSON.parse(s.config), null, 2) // pretty print
            this.setState({snapSettings: snap, settings: settings})
        }
    }

    handleSettingsChange = (e) => {
        e.preventDefault();
        this.setState({settings: e.target.value})
    }

    handleSettingsUpdate = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');

        api.snapsSettingsUpdate(snap, this.state.settings).then(response => {
            this.setState({snapSettings: null, message: 'Sent request to update snap: ' + snap, error: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }

    renderSettings(snap) {
        if (snap.name !== this.state.snapSettings) {
            return
        }

        return (
            <tr>
                <td colSpan="5">
                    <form>
                        <fieldset>
                            <label htmlFor={this.state.snapSettings}>
                                {T('snap-settings')}:
                                <textarea className="col-12" rows="4" value={this.state.settings} data-key={this.state.snapSettings} onChange={this.handleSettingsChange} />
                            </label>
                        </fieldset>
                        <button className="p-button--brand" onClick={this.handleSettingsUpdate} data-key={snap.name}>{T('update')}</button>
                    </form>
                </td>
            </tr>
        )
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
                <h2>{T('snaps-config')}</h2>

                <section className="row">
                    {this.renderMessage()}
                    {this.renderError()}
                    <table className="p-card__content">
                        <thead>
                        <tr>
                            <th className="small">{T('snap')}</th><th className="small">{T('channel')}</th><th>{T('version')}</th><th className="xsmall">{T('status')}</th>
                            <th className="xsmall">{T('settings')}</th>
                        </tr>
                        </thead>
                        {this.state.snaps.map((s) => {
                            var c = '';
                            if (s.name === this.state.snapSettings) {
                                c = 'borderless'
                            }
                            return (
                                <tbody>
                                <tr key={s.name} className={c}>
                                    <td title={s.description}>{s.name}</td>
                                    <td>{s.channel}</td>
                                    <td>{s.version}</td>
                                    <td>{s.status}</td>
                                    <td>
                                        <button className="p-button--neutral small" title={T('view-settings')} data-key={s.name} onClick={this.handleShowSettings}>
                                            <i className="p-icon--menu" aria-hidden="true" data-key={s.name} />
                                        </button>
                                    </td>
                                </tr>
                                {this.renderSettings(s)}
                                </tbody>
                            )
                        })}
                    </table>
                </section>
            </div>
        );
    }
}

export default Snaps;