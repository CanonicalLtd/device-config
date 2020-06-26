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
import FactoryReset from "./FactoryReset";
import ExportImport from "./ExportImport";

class Settings extends Component {
    constructor(props) {
        super(props)
        this.state = {
            selected: 'export-import',
            tabs: ['export-import'],
            error: '',
            message: '',
        };
    }

    componentDidMount() {
        if (this.props.config.factoryReset) {
            this.setState({tabs: ['export-import', 'factory-reset']})
        }
    }

    handleSelectTab = (e) => {
        e.preventDefault()
        let tab = e.target.getAttribute('data-key')
        this.setState({selected: tab, error: '', message:''})
    }

    render() {
        return (
            <div>
                <h2>{T('settings')}</h2>

                <nav className="p-tabs">
                    <ul className="p-tabs__list" role="tablist">
                        {this.state.tabs.map((t) => {
                            let selected = "false"
                            if (this.state.selected === t) {
                                selected = "true"
                            }
                            return (<li className="p-tabs__item" role="presentation">
                                <a href={'#' + t} data-key={t}
                                   onClick={this.handleSelectTab} className="p-tabs__link" role="tab"
                                   aria-controls="section1" aria-selected={selected}>{T(t)}</a>
                            </li>)
                        })}
                    </ul>
                </nav>

                <div>
                    {this.state.selected==='export-import'? <ExportImport /> : ''}
                    {this.state.selected==='factory-reset'? <FactoryReset /> : ''}
                </div>
            </div>
        );
    }
}

export default Settings;