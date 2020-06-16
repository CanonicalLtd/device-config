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
import GaugeChart from 'react-gauge-chart'
import {Row} from "@canonical/react-components";
import {T} from "./Utils";


function chart(id, percent) {
    return (
        <div className="col-2">
            <h5>{T(id)}</h5>
            <GaugeChart id={id}
                        nrOfLevels={420}
                        arcsLength={[0.3, 0.5, 0.2]}
                        colors={['#1c6a08', '#d8d404', '#af1904']}
                        percent={percent/100.0}
                        arcPadding={0.0}
                        textColor={'#2d2d2d'}
                        cornerRadius={0}
            />
        </div>
    )

}

function SystemResources(props) {
    return (
        <Row className="u-align--center">
            <div className="col-3"></div>
            {chart("processor", props.system.cpu)}
            {chart("memory", props.system.memory)}
            {chart("disk", props.system.disk)}
        </Row>
    );
}

export default SystemResources;