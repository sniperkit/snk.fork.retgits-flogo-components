/**
 * Imports
 */
import { Observable } from "rxjs/Observable";
import { Injectable, Injector, Inject } from "@angular/core";
import { Http } from "@angular/http";
import {
    WiContrib,
    WiProxyCORSUtils,
    WiServiceHandlerContribution,
    IValidationResult,
    ValidationResult,
    IFieldDefinition,
    IActivityContribution,
    IConnectorContribution,
    WiContributionUtils
} from "wi-studio/app/contrib/wi-contrib";
import {HTTP_METHOD} from "wi-studio/common/models/contrib";
import { JsonSchema } from "./activity.jsonschema";

var projectState = "", issueState = "", inputState = "", outputState = "";

@WiContrib({})
@Injectable()

export class JiraCreateTicketActivityContribution extends WiServiceHandlerContribution {
    private category: string;
    constructor( @Inject(Injector) injector, private http: Http) {
        super(injector, http);
        this.category = "Jira";
    }

    value = (fieldName: string, context: IActivityContribution): Observable<any> | any => {
        if (fieldName === "Connection") {
            return Observable.create(observer => {
                let connectionRefs = [];

                WiContributionUtils.getConnections(this.http, this.category).subscribe((data: IConnectorContribution[]) => {
                    data.forEach(connection => {
                        for (let i = 0; i < connection.settings.length; i++) {
                            if (connection.settings[i].name === "name") {
                                connectionRefs.push({
                                    "unique_id": WiContributionUtils.getUniqueId(connection),
                                    "name": connection.settings[i].value
                                });
                                break;
                            }
                        }
                    });
                    observer.next(connectionRefs);
                });
            });
        } else if (fieldName === "project") {
            let connectionField: IFieldDefinition = context.getField("Connection");

            if(connectionField.value) {
                let str = connectionField.value
                let isChanged = true;
                if(projectState === "" || projectState != str){
                    projectState = str;
                }else{
                    isChanged = false;
                }
                if(isChanged){
                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, connectionField.value)
                    .map(data => data)
                    .subscribe(data =>{
                        let domain = "", userName = "", password = "";
                        for (let configuration of data.settings) {
                            if (configuration.name === "domain") {
                                domain = configuration.value
                            } else if (configuration.name === "userName") {
                                userName = configuration.value
                            } else if (configuration.name === "password") {
                                password = configuration.value
                            }
                        }

                        let jiraProjectURL = domain + "/rest/api/2/project"
                        console.log("Before Project REST API!!");
                        WiProxyCORSUtils.createRequest(this.http, jiraProjectURL)
                            .addMethod(HTTP_METHOD.GET)
                            .addHeader("Content-Type", "application/json")
                            .addHeader("Authorization", "Basic " + btoa(userName + ":" + password))
                            .send().subscribe(resp => {
                                console.log("Successful result from Project REST API!!");
                                var projectKeys : string[] = new Array();
                                for(let i = 0; i < resp.json().length; i++) {
                                    projectKeys.push(resp.json()[i].key)
                                }
                                projectKeys = projectKeys.sort();
                                observer.next(projectKeys);
                            },
                            error => {
                                console.log("Failed to get fields");
                                observer.next("{}");
                            }
                        );
                    });
                });
            }
            }
        } else if (fieldName === "issueType") {
            let connectionField: IFieldDefinition = context.getField("Connection");
            let projectField: IFieldDefinition = context.getField("project");

            if(connectionField.value && projectField.value) {
                let str = connectionField.value + projectField.value
                let isChanged = true;
                if(issueState === "" || issueState != str){
                    issueState = str;
                }else{
                    isChanged = false;
                }
                if(isChanged){
                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, connectionField.value)
                    .map(data => data)
                    .subscribe(data =>{
                        let domain = "", userName = "", password = "";
                        for (let configuration of data.settings) {
                            if (configuration.name === "domain") {
                                domain = configuration.value
                            } else if (configuration.name === "userName") {
                                userName = configuration.value
                            } else if (configuration.name === "password") {
                                password = configuration.value
                            }
                        }

                        let jiraIssueTypeURL = domain + "/rest/api/2/project/" + projectField.value
                        console.log("Before IssueType REST API!!");
                        WiProxyCORSUtils.createRequest(this.http, jiraIssueTypeURL)
                            .addMethod(HTTP_METHOD.GET)
                            .addHeader("Content-Type", "application/json")
                            .addHeader("Authorization", "Basic " + btoa(userName + ":" + password))
                            .send().subscribe(resp => {
                                console.log("Successful result from IssueType REST API!!");
                                var issueTypes : string[] = new Array();
                                for(let i = 0; i < resp.json().issueTypes.length; i++) {
                                    issueTypes.push(resp.json().issueTypes[i].name)
                                }
                                issueTypes = issueTypes.sort();
                                observer.next(issueTypes);
                            },
                            error => {
                                console.log("Failed to get Issue Types");
                                observer.next("{}");
                            }
                        );
                    });
                });
            }
            }
        } else if (fieldName === "input") {
            let connectionField: IFieldDefinition = context.getField("Connection");
            let projectField: IFieldDefinition = context.getField("project");
            let issueTypeField: IFieldDefinition = context.getField("issueType");

            if(connectionField.value && projectField.value && issueTypeField.value){
                let str = connectionField.value + projectField.value + issueTypeField.value
                let isChanged = true;
                if(inputState === "" || inputState != str){
                    inputState = str;
                }else{
                    isChanged = false;
                }
                if(isChanged) {
                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, connectionField.value)
                    .map(data => data)
                    .subscribe(data =>{
                        let domain = "", userName = "", password = "";
                        for (let configuration of data.settings) {
                            if (configuration.name === "domain") {
                                domain = configuration.value
                            } else if (configuration.name === "userName") {
                                userName = configuration.value
                            } else if (configuration.name === "password") {
                                password = configuration.value
                            }
                        }

                        let jiraMetadataURL = domain + "/rest/api/2/issue/createmeta?projectKeys="+projectField.value+"&issuetypeNames="+issueTypeField.value+"&expand=projects.issuetypes.fields";
                        console.log("Before Input Metadata REST API!!")
                        WiProxyCORSUtils.createRequest(this.http, jiraMetadataURL)
                            .addMethod(HTTP_METHOD.GET)
                            .addHeader("Content-Type", "application/json")
                            .addHeader("Authorization", "Basic " + btoa(userName + ":" + password))
                            .send().subscribe(resp => {
                                console.log("Successful result from Input Metadata REST API!!");
                                let jsonRespFields = resp.json().projects[0].issuetypes[0].fields;
                                let inputSchema = this.getFields(jsonRespFields,fieldName)
                                observer.next(JSON.stringify(inputSchema));
                            },
                            error => {
                                console.log("Failed to get fields");
                                observer.next("{}");
                            });
                        });
                    });
                 }
            }
        } else if (fieldName === "output") {
            let connectionField: IFieldDefinition = context.getField("Connection");
            let projectField: IFieldDefinition = context.getField("project");
            let issueTypeField: IFieldDefinition = context.getField("issueType");

            if(connectionField.value && projectField.value && issueTypeField.value){
                let str = connectionField.value + projectField.value + issueTypeField.value
                let isChanged = true;
                if(outputState === "" || outputState != str){
                    outputState = str;
                }else{
                    isChanged = false;
                }
                if(isChanged) {
                return Observable.create(observer => {
                    WiContributionUtils.getConnection(this.http, connectionField.value)
                    .map(data => data)
                    .subscribe(data =>{
                        let domain = "", userName = "", password = "";
                        for (let configuration of data.settings) {
                            if (configuration.name === "domain") {
                                domain = configuration.value
                            } else if (configuration.name === "userName") {
                                userName = configuration.value
                            } else if (configuration.name === "password") {
                                password = configuration.value
                            }
                        }

                        let jiraMetadataURL = domain + "/rest/api/2/issue/createmeta?projectKeys="+projectField.value+"&issuetypeNames="+issueTypeField.value+"&expand=projects.issuetypes.fields";
                        console.log("Before Output Metadata REST API!!")
                        WiProxyCORSUtils.createRequest(this.http, jiraMetadataURL)
                            .addMethod(HTTP_METHOD.GET)
                            .addHeader("Content-Type", "application/json")
                            .addHeader("Authorization", "Basic " + btoa(userName + ":" + password))
                            .send().subscribe(resp => {
                                console.log("Successful result from Output Metadata REST API!!");
                                let jsonRespFields = resp.json().projects[0].issuetypes[0].fields;
                                let outputSchema = this.getFields(jsonRespFields,fieldName)
                                observer.next(JSON.stringify(outputSchema));
                            },
                            error => {
                                console.log("Failed to get fields");
                                observer.next("{}");
                            });
                        });
                    });
                }
            }
        }
        return null;
    }
    validate = (fieldName: string, context: IActivityContribution): Observable<IValidationResult> | IValidationResult => {
        if (fieldName === "Connection") {
            let connection: IFieldDefinition = context.getField("Connection")
            if (connection.value === null) {
                return ValidationResult.newValidationResult().setError("JIRA-1000", "Jira Connection must be configured");
            }
        }
        return null;
    }

    getFields = (jsonRespFields : any, fieldName: string) : any => {
        let notAllowedStandardFields: string[] = ["issuetype","project","issuelinks","timetracking","attachment"];
        let notAllowedCustomType : string[] = ["com.pyxis.greenhopper.jira:gh-sprint","com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect"];
        let reqFields = {}, optFields = {};
        var requiredFields : string[] = new Array();
        for (let jiraFieldName in jsonRespFields) {
            if(notAllowedStandardFields.indexOf(jiraFieldName) < 0 && notAllowedCustomType.indexOf(jsonRespFields[jiraFieldName].schema.custom) < 0) {
                if(jsonRespFields[jiraFieldName].required){
                    reqFields[jsonRespFields[jiraFieldName].name] = JsonSchema.Types.toJsonType(jsonRespFields[jiraFieldName].schema, jiraFieldName )
                    requiredFields.push(jsonRespFields[jiraFieldName].name)
                } else {
                    optFields[jsonRespFields[jiraFieldName].name] = JsonSchema.Types.toJsonType(jsonRespFields[jiraFieldName].schema, jiraFieldName)
                }
            }
        }
        var fields = { ...reqFields, ...optFields };
        if (fieldName === "output") {
            let outputKey = { "IssueID" : {
                "type":"string"
            }}
            fields = { ...outputKey, ...fields}
        }

        let schema = {
            "properties": fields, "required": requiredFields, "type": "object"
        };
        return schema;
    }
}