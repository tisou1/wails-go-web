export namespace main {
	
	export class DialogOption {
	    title: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new DialogOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.title = source["title"];
	        this.message = source["message"];
	    }
	}

}

