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

import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import App from './App';
import {T, getAppConfig} from "./components/Utils";


getAppConfig( (cfg) => {
    if (!cfg.custom) {
        cfg.custom = {
            copyright: T('copyright'), title: T('title'), subtitle: T('subtitle'),
            bullets: [
                {text: T('legal'), url: 'https://ubuntu.com/legal'},
                {text: T('privacy'), url: 'https://ubuntu.com/legal/data-privacy'},
                {text: T('report-a-bug'), url: 'https://github.com/CanonicalLtd/device-config/issues/new'},
            ]
        }
    }
    ReactDOM.render(<App config={cfg} />, document.getElementById('root'));
})