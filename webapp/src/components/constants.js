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

const API_PREFIX = '/v1/'

function getBaseURL() {
    return window.location.protocol + '//' + window.location.hostname + ':' + window.location.port + API_PREFIX;
}

let Constants = {
    baseUrl: getBaseURL(),
    LoadingImage: '/static/images/ajax-loader.gif',
    languages: ['de', 'en','sv'],
    //missingIcon: '/static/images/snapcraft-missing-icon.svg',
}

export default Constants
