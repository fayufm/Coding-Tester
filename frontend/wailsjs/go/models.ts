export namespace main {
	
	export class AIConfig {
	    selectedProviderId: string;
	    apiKey: string;
	    customEndpoint: string;
	
	    static createFrom(source: any = {}) {
	        return new AIConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.selectedProviderId = source["selectedProviderId"];
	        this.apiKey = source["apiKey"];
	        this.customEndpoint = source["customEndpoint"];
	    }
	}
	export class AIProvider {
	    id: string;
	    name: string;
	    endpointUrl: string;
	    documentUrl: string;
	    apiKeyHeader: string;
	    apiKeyPrefix: string;
	
	    static createFrom(source: any = {}) {
	        return new AIProvider(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.endpointUrl = source["endpointUrl"];
	        this.documentUrl = source["documentUrl"];
	        this.apiKeyHeader = source["apiKeyHeader"];
	        this.apiKeyPrefix = source["apiKeyPrefix"];
	    }
	}
	export class AIResponse {
	    content: string;
	    error: string;
	    success: boolean;
	    provider: string;
	
	    static createFrom(source: any = {}) {
	        return new AIResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.error = source["error"];
	        this.success = source["success"];
	        this.provider = source["provider"];
	    }
	}
	export class LanguageConfig {
	    language: string;
	
	    static createFrom(source: any = {}) {
	        return new LanguageConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.language = source["language"];
	    }
	}
	export class PackageInfo {
	    name: string;
	    version: string;
	    description: string;
	    installed: boolean;
	    installLink: string;
	    downloadUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new PackageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.version = source["version"];
	        this.description = source["description"];
	        this.installed = source["installed"];
	        this.installLink = source["installLink"];
	        this.downloadUrl = source["downloadUrl"];
	    }
	}
	export class LanguageInfo {
	    name: string;
	    installed: boolean;
	    version: string;
	    missingDeps: string[];
	    downloadUrl: string;
	    installTutorial: string;
	    packageManager: string;
	    packages: PackageInfo[];
	    extensions: PackageInfo[];
	    recommendedPkgs: PackageInfo[];
	
	    static createFrom(source: any = {}) {
	        return new LanguageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.installed = source["installed"];
	        this.version = source["version"];
	        this.missingDeps = source["missingDeps"];
	        this.downloadUrl = source["downloadUrl"];
	        this.installTutorial = source["installTutorial"];
	        this.packageManager = source["packageManager"];
	        this.packages = this.convertValues(source["packages"], PackageInfo);
	        this.extensions = this.convertValues(source["extensions"], PackageInfo);
	        this.recommendedPkgs = this.convertValues(source["recommendedPkgs"], PackageInfo);
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
	
	export class PackageTutorial {
	    name: string;
	    installCmd: string;
	    searchCmd: string;
	    updateCmd: string;
	    tutorialUrl: string;
	
	    static createFrom(source: any = {}) {
	        return new PackageTutorial(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.installCmd = source["installCmd"];
	        this.searchCmd = source["searchCmd"];
	        this.updateCmd = source["updateCmd"];
	        this.tutorialUrl = source["tutorialUrl"];
	    }
	}
	export class SystemInfo {
	    os: string;
	    arch: string;
	    cpus: string;
	
	    static createFrom(source: any = {}) {
	        return new SystemInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.os = source["os"];
	        this.arch = source["arch"];
	        this.cpus = source["cpus"];
	    }
	}
	export class ThemeConfig {
	    theme: string;
	
	    static createFrom(source: any = {}) {
	        return new ThemeConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	    }
	}

}

