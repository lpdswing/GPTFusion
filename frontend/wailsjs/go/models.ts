export namespace main {
	
	export class PlatForm {
	    id: string;
	    label: string;
	    url: string;
	    priority: number;
	    separator: boolean;
	    group: string;
	
	    static createFrom(source: any = {}) {
	        return new PlatForm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.label = source["label"];
	        this.url = source["url"];
	        this.priority = source["priority"];
	        this.separator = source["separator"];
	        this.group = source["group"];
	    }
	}
	export class Setting {
	    mode: string;
	    always_on_top: boolean;
	    hide_window_on_close: boolean;
	    remember_last_page: boolean;
	    last_page: string;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mode = source["mode"];
	        this.always_on_top = source["always_on_top"];
	        this.hide_window_on_close = source["hide_window_on_close"];
	        this.remember_last_page = source["remember_last_page"];
	        this.last_page = source["last_page"];
	    }
	}

}

