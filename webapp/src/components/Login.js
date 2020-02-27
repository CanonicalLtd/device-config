/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

import React, { Component } from 'react';
import api from './api';
import {T, setSession} from './Utils';
import AlertBox from './AlertBox';

class Login extends Component {
    constructor(props) {
        super(props)
        this.state = {
            macAddress: '',
            error: '',
        };
    }

    handleMacAddressChange = (e) => {
        e.preventDefault()
        this.setState({macAddress: e.target.value})
    }

    login() {
        api.loginRequest({macAddress: this.state.macAddress}).then(response => {
            if (!response.data.code) {
                setSession(response.data.username, response.data.sessionId)
                window.location.href = "/"
                return
            }

            this.setState({error: T(response.data.code) + ": " + response.data.message})
        }).catch(e => {
            console.log(e.response)
            this.setState({error: T(e.response.data.code) + ": " + e.response.data.message})
        })
    }

    handleSubmit = (e) => {
        e.preventDefault()
        this.login()
    }

    renderError() {
        if (this.state.error) {
            return (
                <AlertBox message={this.state.error} />
            );
        }
    }

    render() {
        return (
            <div>
                <div className="row">
                    <h2>{T("login")}</h2>
                    <p>{T("login-description")}</p>
                </div>

                {this.renderError()}
                <form>
                    <label for="macaddress">MAC Address:</label>
                    <input name="macaddress" type="text" value={this.state.macAddress} onChange={this.handleMacAddressChange} />
                    <button className="p-button--positive" onClick={this.handleSubmit}>{T("submit")}</button>
                </form>

            </div>
        );
    }
}

export default Login;