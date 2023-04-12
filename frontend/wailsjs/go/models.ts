export namespace main {
	
	export class PlatForm {
	    id: string;
	    label: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new PlatForm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.url = source["url"];
	    }
	}

}

