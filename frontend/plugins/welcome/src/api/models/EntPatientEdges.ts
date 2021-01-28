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
    EntBloodtype,
    EntBloodtypeFromJSON,
    EntBloodtypeFromJSONTyped,
    EntBloodtypeToJSON,
    EntCategory,
    EntCategoryFromJSON,
    EntCategoryFromJSONTyped,
    EntCategoryToJSON,
    EntDiagnosis,
    EntDiagnosisFromJSON,
    EntDiagnosisFromJSONTyped,
    EntDiagnosisToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntNametitle,
    EntNametitleFromJSON,
    EntNametitleFromJSONTyped,
    EntNametitleToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * 
     * @type {EntBloodtype}
     * @memberof EntPatientEdges
     */
    bloodtype?: EntBloodtype;
    /**
     * 
     * @type {EntCategory}
     * @memberof EntPatientEdges
     */
    category?: EntCategory;
    /**
     * Diagnosis holds the value of the diagnosis edge.
     * @type {Array<EntDiagnosis>}
     * @memberof EntPatientEdges
     */
    diagnosis?: Array<EntDiagnosis>;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntPatientEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntNametitle}
     * @memberof EntPatientEdges
     */
    nametitle?: EntNametitle;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bloodtype': !exists(json, 'bloodtype') ? undefined : EntBloodtypeFromJSON(json['bloodtype']),
        'category': !exists(json, 'category') ? undefined : EntCategoryFromJSON(json['category']),
        'diagnosis': !exists(json, 'diagnosis') ? undefined : ((json['diagnosis'] as Array<any>).map(EntDiagnosisFromJSON)),
        'employee': !exists(json, 'employee') ? undefined : EntEmployeeFromJSON(json['employee']),
        'gender': !exists(json, 'gender') ? undefined : EntGenderFromJSON(json['gender']),
        'nametitle': !exists(json, 'nametitle') ? undefined : EntNametitleFromJSON(json['nametitle']),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bloodtype': EntBloodtypeToJSON(value.bloodtype),
        'category': EntCategoryToJSON(value.category),
        'diagnosis': value.diagnosis === undefined ? undefined : ((value.diagnosis as Array<any>).map(EntDiagnosisToJSON)),
        'employee': EntEmployeeToJSON(value.employee),
        'gender': EntGenderToJSON(value.gender),
        'nametitle': EntNametitleToJSON(value.nametitle),
    };
}


