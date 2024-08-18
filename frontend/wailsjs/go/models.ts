export namespace assetmanager {
	
	export class Asset {
	    uuid: string;
	    name: string;
	    lastUpdate: string;
	
	    static createFrom(source: any = {}) {
	        return new Asset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uuid = source["uuid"];
	        this.name = source["name"];
	        this.lastUpdate = source["lastUpdate"];
	    }
	}
	export class AssetFile {
	    uuid: string;
	    name: string;
	    description: string;
	    extension: string;
	
	    static createFrom(source: any = {}) {
	        return new AssetFile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uuid = source["uuid"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.extension = source["extension"];
	    }
	}
	export class CreateAssetFilePayload {
	    assetUuid: string;
	    name: string;
	    description: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateAssetFilePayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetUuid = source["assetUuid"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.path = source["path"];
	    }
	}
	export class CreateAssetFileResponse {
	    assetFile: AssetFile;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateAssetFileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetFile = this.convertValues(source["assetFile"], AssetFile);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CreateAssetResponse {
	    id: string;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateAssetResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.error = source["error"];
	    }
	}
	export class DeleteAssetFilePayload {
	    assetUuid: string;
	    fileUuid: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteAssetFilePayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetUuid = source["assetUuid"];
	        this.fileUuid = source["fileUuid"];
	    }
	}
	export class DeleteAssetFileResponse {
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteAssetFileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	    }
	}
	export class EditAssetFilePayload {
	    assetUuid: string;
	    fileUuid: string;
	    name: string;
	    description: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new EditAssetFilePayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetUuid = source["assetUuid"];
	        this.fileUuid = source["fileUuid"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.path = source["path"];
	    }
	}
	export class EditAssetFileResponse {
	    assetFile: AssetFile;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new EditAssetFileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetFile = this.convertValues(source["assetFile"], AssetFile);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetAssetResponse {
	    asset: Asset;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new GetAssetResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.asset = this.convertValues(source["asset"], Asset);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ListAssetFilesResponse {
	    assetFiles: AssetFile[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ListAssetFilesResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assetFiles = this.convertValues(source["assetFiles"], AssetFile);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ListAssetsResponse {
	    assets: Asset[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new ListAssetsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assets = this.convertValues(source["assets"], Asset);
	        this.error = source["error"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace dbmanager {
	
	export class DatabaseManager {
	
	
	    static createFrom(source: any = {}) {
	        return new DatabaseManager(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace main {
	
	export class SelectFileResponse {
	    path: string;
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new SelectFileResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.error = source["error"];
	    }
	}

}

export namespace tagmanager {
	
	export class TagSearchPayload {
	    search: string;
	    omit: string[];
	
	    static createFrom(source: any = {}) {
	        return new TagSearchPayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.search = source["search"];
	        this.omit = source["omit"];
	    }
	}
	export class TagSearchResponse {
	    tags: string[];
	    error: string;
	
	    static createFrom(source: any = {}) {
	        return new TagSearchResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tags = source["tags"];
	        this.error = source["error"];
	    }
	}

}

export namespace vault {
	
	export class Vault {
	
	
	    static createFrom(source: any = {}) {
	        return new Vault(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

