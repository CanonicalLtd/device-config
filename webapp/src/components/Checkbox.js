/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
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