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
	export class Setting {
	    mode: string;
	    always_on_top: boolean;
	    hide_window_on_close: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mode = source["mode"];
	        this.always_on_top = source["always_on_top"];
	        this.hide_window_on_close = source["hide_window_on_close"];
	    }
	}

}

