resource "null_resource" "get_user" {
    triggers  =  { always_run = "${timestamp()}" }
    provisioner "local-exec" {
        command = "whoami | awk '{print \"User: \"$0\"\"}' > ${path.module}/local-exec-out"
    }
}

resource "null_resource" "get_timestamp" {
    triggers  =  { always_run = "${timestamp()}" }
    provisioner "local-exec" {
        command = "date | awk '{print \"Date: \"$0\"\"}' >> ${path.module}/local-exec-out"
    }
    depends_on = [
        null_resource.get_user
    ]
}

/*
resource "local_file" "write_demo_file" {
  content  = data.template_file.demo_template.rendered
  filename = "${path.module}/demo-output.txt"
}
*/