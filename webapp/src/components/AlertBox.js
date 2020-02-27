/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

import React, {Component} from 'react';


class AlertBox extends Component {
    render() {
        if (this.props.message) {
            var c = 'p-notification--';
            if (this.props.type) {
                c = c + this.props.type;
            } else {
                c = c + 'negative';
            }

            return (
                <div className={c}>
                    <p className="p-notification__response">{this.props.message}</p>
                </div>
            );
        } else {
            return <span />;
        }
    }
}

export default AlertBox;