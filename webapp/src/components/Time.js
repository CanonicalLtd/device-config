/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */


import React, {Component} from 'react';
import api from "./api";
import {formatError, T} from "./Utils";
import AlertBox from "./AlertBox";
import Checkbox from "./Checkbox";

class Time extends Component {
    constructor(props) {
        super(props)
        this.state = {
            time: {timezones: []},
            error: '',
            message: '',
        };
    }

    componentDidMount() {
        this.getTimeConfig()
    }

    getTimeConfig = () => {
        api.timeGet().then(response => {
            this.setState({time: response.data.time, error: ''})
        })
        .catch(e => {
            this.setState({message: formatError(e.response)});
        })
    }

    setField(field, value) {
        let t = this.state.time
        t[field] = value
        this.setState({time: t})
    }

    handleTimezoneChange = (e) => {
        e.preventDefault()
        this.setField('timezone', e.target.value)
    }

    handleNTPChange = (value) => {
        this.setField('ntp', value)
    }

    handleTimeChange = (e) => {
        e.preventDefault()
        this.setField('time', e.target.value)
    }

    handleSave = (e) => {
        e.preventDefault()

        // Save the time config
        api.timeUpdate(this.state.time).then(response => {
            this.getTimeConfig()
            this.setState({message: T('time-updated'), error: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response), message: ''});
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
                <h2>{T('time-config')}</h2>

                <section className="row">
                    {this.renderMessage()}
                    {this.renderError()}
                    <form>
                        <fieldset>
                            <label htmlFor={"timezone"}>{T('timezone')}:</label>
                            <select value={this.state.time.timezone} onChange={this.handleTimezoneChange} placeholder={T('timezone-help')}>
                                <option/>
                                {this.state.time.timezones.map((m) => {
                                    return <option value={m}>{T(m)}</option>
                                })}
                            </select>

                            <Checkbox label={T('ntp')} checked={this.state.time.ntp} onChange={this.handleNTPChange} />
                            <label htmlFor={"time"}>{T('time')}:</label>
                            <input name="time" type="text" value={this.state.time.time} onChange={this.handleTimeChange}
                                   placeholder={T('time-help')} disabled={this.state.time.ntp}/>

                            <button onClick={this.handleSave} className="p-button--positive">{T('save')}</button>
                        </fieldset>
                    </form>
                </section>
            </div>
        );
    }
}


export default Time;
