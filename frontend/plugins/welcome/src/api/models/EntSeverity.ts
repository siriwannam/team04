/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntSeverityEdges,
    EntSeverityEdgesFromJSON,
    EntSeverityEdgesFromJSONTyped,
    EntSeverityEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSeverity
 */
export interface EntSeverity {
    /**
     * SeverityName holds the value of the "SeverityName" field.
     * @type {string}
     * @memberof EntSeverity
     */
    severityName?: string;
    /**
     * 
     * @type {EntSeverityEdges}
     * @memberof EntSeverity
     */
    edges?: EntSeverityEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSeverity
     */
    id?: number;
}

export function EntSeverityFromJSON(json: any): EntSeverity {
    return EntSeverityFromJSONTyped(json, false);
}

export function EntSeverityFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSeverity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'severityName': !exists(json, 'SeverityName') ? undefined : json['SeverityName'],
        'edges': !exists(json, 'edges') ? undefined : EntSeverityEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSeverityToJSON(value?: EntSeverity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'SeverityName': value.severityName,
        'edges': EntSeverityEdgesToJSON(value.edges),
        'id': value.id,
    };
}

