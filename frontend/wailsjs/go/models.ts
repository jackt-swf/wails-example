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

}

