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

import Messages from './Messages'
import Cookies from 'js-cookie'

export function T(message) {
    const msg = Messages[message] || message;
    return msg
}

// URL is in the form:
//  /section
//  /section/sectionId
//  /section/sectionId/subsection
export function parseRoute() {
    const parts = window.location.pathname.split('/')

    switch (parts.length) {
        case 2:
            return {section: parts[1]}
        case 3:
            return {section: parts[1], sectionId: parts[2]}
        case 4:
            return {section: parts[1], sectionId: parts[2], subsection: parts[3]}
        default:
            return {}
    }
}

export function checkSession() {
    let username = Cookies.get('username')
    let sessionId = Cookies.get('sessionID')

    console.log("Username/Session:", username, sessionId)

    if ((username) && (sessionId)) {
        return true
    }
    return false
}

export function setSession(u, s) {
    Cookies.set('username', u)
    Cookies.set('sessionID', s)
}

export function formatError(data) {
    let message = T(data.code);
    if (data.message) {
        message += ': ' + data.message;
    }
    return message;
}
