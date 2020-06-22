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

import axios from 'axios'
import constants from './constants'

let service = {
    configGet:  (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'config');
    },

    loginRequest:  (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'login', query);
    },

    factoryReset:  (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'factory-reset', query);
    },

    networkGet: () => {
        return axios.get(constants.baseUrl + 'network');
    },

    networkUpdate: (iface) => {
        return axios.post(constants.baseUrl + 'network', iface);
    },

    networkApply: () => {
        return axios.post(constants.baseUrl + 'network/apply');
    },

    proxyGet: () => {
        return axios.get(constants.baseUrl + 'proxy');
    },

    proxyUpdate: (iface) => {
        return axios.post(constants.baseUrl + 'proxy', iface);
    },

    timeGet: () => {
        return axios.get(constants.baseUrl + 'time');
    },

    timeUpdate: (t) => {
        return axios.post(constants.baseUrl + 'time', {ntp: t.ntp, time: t.time, timezone: t.timezone});
    },

    servicesGet: () => {
        return axios.get(constants.baseUrl + 'services');
    },

    systemResourcesGet: () => {
        return axios.get(constants.baseUrl + 'system');
    },

    snapsGet: () => {
        return axios.get(constants.baseUrl + 'snaps');
    },

    snapsSettingsUpdate: (snap, settings) => {
        return axios.put(constants.baseUrl + 'snaps/' + snap, settings);
    },

}

export default service