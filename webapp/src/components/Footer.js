/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

import React, {Component} from 'react';
import {T} from './Utils'

class Footer extends Component {
    render() {
        return (
            <div id="footer">
                <div className="row footer">
                    <p>{T('copyright')}</p>
                    <nav className="p-footer__nav" role="navigation">
                        <ul className="p-footer__links">
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://ubuntu.com/legal">{T('legal')}</a>
                            </li>
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://ubuntu.com/legal/data-privacy">{T('privacy')}</a>
                            </li>
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://github.com/CanonicalLtd/imagebuild/issues/new">{T('report-a-bug')}</a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
        );
    }
}

export default Footer;