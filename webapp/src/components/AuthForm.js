import React from 'react';
import {T} from "./Utils";
import {Form, Input, Button} from "@canonical/react-components";

function AuthForm(props) {
    return (
        <Form>
            <label htmlFor="macaddress">{T('macaddress')}:</label>
            <Input id="macaddress" type="text" value={props.macAddress} onChange={props.onChange} />
            <Button appearance="positive" onClick={props.onClick}>{T("submit")}</Button>
        </Form>
    );
}

export default AuthForm;