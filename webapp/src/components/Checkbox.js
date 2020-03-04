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
import {T} from "./Utils";

class Checkbox extends Component {
    handleUnchecked  = (e) => {
        e.preventDefault()
        this.props.onChange(false)
    }

    handleChecked  = (e) => {
        e.preventDefault()
        this.props.onChange(true)
    }

    render() {
        if (this.props.checked) {
            return (
                <div>
                    <a href="#select" className="p-button--base has-icon" onClick={this.handleUnchecked}><img src="/static/images/checkbox_checked_16.png" alt="checked" /></a>
                    <span>{T(this.props.label)}</span>
                </div>
            )
        } else {
            return (
                <div>
                    <a href="#select" className="p-button--base has-icon" onClick={this.handleChecked}><img src="/static/images/checkbox_unchecked_16.png" alt="unchecked"  /></a>
                    <span>{T(this.props.label)}</span>
                </div>
            )
        }
    }
}

export default Checkbox;