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
    EntDisease,
    EntDiseaseFromJSON,
    EntDiseaseFromJSONTyped,
    EntDiseaseToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDiagnosisEdges
 */
export interface EntDiagnosisEdges {
    /**
     * 
     * @type {EntDisease}
     * @memberof EntDiagnosisEdges
     */
    disease?: EntDisease;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntDiagnosisEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntDiagnosisEdges
     */
    patient?: EntPatient;
}

export function EntDiagnosisEdgesFromJSON(json: any): EntDiagnosisEdges {
    return EntDiagnosisEdgesFromJSONTyped(json, false);
}

export function EntDiagnosisEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDiagnosisEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'disease': !exists(json, 'Disease') ? undefined : EntDiseaseFromJSON(json['Disease']),
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'patient': !exists(json, 'Patient') ? undefined : EntPatientFromJSON(json['Patient']),
    };
}

export function EntDiagnosisEdgesToJSON(value?: EntDiagnosisEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'disease': EntDiseaseToJSON(value.disease),
        'employee': EntEmployeeToJSON(value.employee),
        'patient': EntPatientToJSON(value.patient),
    };
}


