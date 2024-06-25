export namespace main {
	
	export class ComputerInformation {
	    hostname: string;
	    macAddress: string;
	    user: string;
	
	    static createFrom(source: any = {}) {
	        return new ComputerInformation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.macAddress = source["macAddress"];
	        this.user = source["user"];
	    }
	}
	export class Config {
	    tenantUrl: string;
	    userEmail: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tenantUrl = source["tenantUrl"];
	        this.userEmail = source["userEmail"];
	    }
	}

}

