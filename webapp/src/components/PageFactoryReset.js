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
import AuthForm from "./AuthForm";
import AlertBox from "./AlertBox";
import api from "./api";

class FactoryReset extends Component {
    constructor(props) {
        super(props)
        this.state = {
            macAddress: '',
            error: '',
        };
    }

    reset() {
        api.factoryReset({macAddress: this.state.macAddress}).then(response => {
            // Will trigger a reboot
        }).catch(e => {
            console.log(e.response)
            this.setState({error: T(e.response.data.code) + ": " + e.response.data.message})
        })
    }

    handleMacAddressChange = (e) => {
        e.preventDefault()
        this.setState({macAddress: e.target.value})
    }

    handleSubmit = (e) => {
        e.preventDefault()
        this.reset()
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
                <h2>{T('factory-reset')}</h2>
                <p>{T('factory-reset-desc')}</p>

                {this.renderError()}
                <AuthForm macAddress={this.state.macAddress} onChange={this.handleMacAddressChange} onClick={this.handleSubmit} />
            </div>
        );
    }
}

export default FactoryReset;