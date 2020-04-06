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

import React, { Component } from 'react';
import {formatError, T} from './Utils';
import AlertBox from './AlertBox';
import api from "./api"

const methods = ['dhcp', 'manual']

class Network extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: null,
            interfaces: [],
            interface: {},
            error: '',
            message: '',
        };
    }

    componentDidMount() {
        this.getNetworkConfig()
    }

    getNetworkConfig = () => {
        api.networkGet().then(response => {
            // Skip the hidden interfaces
            let interfaces = response.data.interfaces.filter(iface => {
                return !this.props.config.hideInterfaces.includes(iface.interface)
            })

            this.setState({interfaces: interfaces, error: ''})

            if ((!this.state.selected) && (response.data.interfaces.length > 0)) {
                this.getInterface(interfaces[0].interface)
                this.setState({selected: interfaces[0].interface})
            }
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data)});
        })
    }

    handleSelectTab = (e) => {
        e.preventDefault()
        let iface = e.target.getAttribute('data-key')
        this.getInterface(iface)
        this.setState({selected: iface, error: '', message:''})
    }

    handleUseChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.use = !iface.use
        this.setState({interface: iface})
    }

    handleMethodChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.method = e.target.value
        this.setState({interface: iface})
    }

    handleGatewayChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.gateway = e.target.value
        this.setState({interface: iface})
    }

    handleAddressChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.address = e.target.value
        this.setState({interface: iface})
    }

    handleMaskChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.mask = e.target.value
        this.setState({interface: iface})
    }

    handleDNSChange = (e) => {
        e.preventDefault()
        let iface = this.state.interface
        iface.nameServers = e.target.value.split(',')
        this.setState({interface: iface})
    }

    handleSave = (e) => {
        e.preventDefault()

        // Save the interface config
        api.networkUpdate(this.state.interface).then(response => {
            this.getNetworkConfig()
            this.setState({message: T('interface-updated'), error: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }

    handleApply = (e) => {
        e.preventDefault()

        // Apply the config
        api.networkApply().then(response => {
            this.getNetworkConfig()
            this.setState({message: T('interface-updated'), error: ''})
        })
        .catch(e => {
            this.setState({error: formatError(e.response.data), message: ''});
        })
    }

    getInterface(iface) {
        let matches = this.state.interfaces.filter(i => {
            return i.interface===iface
        })
        if (matches.length>0) {
            this.setState({interface: matches[0]})
        }
    }

    renderUse() {
        if (this.state.interface.use) {
            return (
                <div>
                    <a href="#use" className="p-button--base has-icon" onClick={this.handleUseChange}><img src="/static/images/checkbox_checked_16.png" alt="checked" /></a>
                    <span>{T('use')}</span>
                </div>
            )
        } else {
            return (
                <div>
                    <a href="#use" className="p-button--base has-icon" onClick={this.handleUseChange}><img src="/static/images/checkbox_unchecked_16.png" alt="unchecked"  /></a>
                    <span>{T('use')}</span>
                </div>
            )
        }
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
                <h2>{T('network-config')}</h2>

                <section className="row">
                    <div>
                        <nav className="p-tabs">
                            <ul className="p-tabs__list" role="tablist">
                            {this.state.interfaces.map((iface) => {
                                let selected = "false"
                                if (this.state.selected===iface.interface) {
                                    selected = "true"
                                }
                                return (<li className="p-tabs__item" role="presentation">
                                    <a href={'#'+iface.interface} data-key={iface.interface} onClick={this.handleSelectTab} class="p-tabs__link" tabindex="0" role="tab" aria-controls="section1" aria-selected={selected}>{iface.interface}</a>
                                </li>)
                            })}
                            </ul>
                        </nav>
                    </div>
                    <div>
                        {this.renderMessage()}
                        {this.renderError()}
                        {this.state.interface.interface ?
                            <form>
                                {this.renderUse()}
                                <fieldset disabled={!this.state.interface.use}>
                                    <h3>{T('interface')} {this.state.interface.interface}</h3>
                                    <label htmlFor={"method"}>{T('method')}:</label>
                                    <select value={this.state.interface.method} onChange={this.handleMethodChange}>
                                        <option/>
                                        {methods.map((m) => {
                                            return <option value={m}>{T(m)}</option>
                                        })}
                                    </select>
                                    <label htmlFor={"nameServers"}>{T('dns')}:</label>
                                    <input name="nameServers" type="text" onChange={this.handleDNSChange}
                                           value={this.state.interface.nameServers ? this.state.interface.nameServers.toString() : ''} placeholder={T('dns-help')} />
                                    <label htmlFor={"address"}>{T('address')}:</label>
                                    <input name="address" type="text" value={this.state.interface.address} placeholder={T('address-help')} onChange={this.handleAddressChange}/>
                                    <label htmlFor={"mask"}>{T('mask')}:</label>
                                    <input name="mask" type="text" value={this.state.interface.mask} onChange={this.handleMaskChange} placeholder={T('mask-help')}/>
                                    <label for={"gateway"}>{T('gateway')}:</label>
                                    <input name="gateway" type="text" value={this.state.interface.gateway} onChange={this.handleGatewayChange}/>
                                </fieldset>
                                <button onClick={this.handleSave} className="p-button--positive">{T('save')}</button>
                            </form>
                            :
                            <h4>{T('select-interface')}</h4>
                        }
                    </div>
                    <div className="row">
                        <p className="col-10">{T('apply-desc')}</p>
                        <button onClick={this.handleApply} className="p-button--brand col-2">{T('apply')}</button>
                    </div>
                </section>
            </div>
        )
    }

}

export default Network;