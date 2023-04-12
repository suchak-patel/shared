data "template_file" "demo_template" {
    template = file("./files/demo-file.tpl")
    vars = {
        timestamp   = "${yamldecode(data.local_file.write_demo_file.content)["Date"]}"
        user        = "${yamldecode(data.local_file.write_demo_file.content)["User"]}"
        tools       = var.tools
    }
    depends_on = [
        null_resource.get_timestamp
    ]
}

data "local_file" "write_demo_file" {
    filename = "${path.module}/local-exec-out"
    depends_on = [
        null_resource.get_timestamp
    ]
}
