document.getElementById("accountForm").addEventListener("submit", async function(event) {
    event.preventDefault();
    const form = event.target;

    const formData = new FormData();
    formData.append("name", form.name.value);
    formData.append("email", form.email.value);

    let formID;
    try {
        const response = await fetch("/submit-form", {
            method: "POST",
            body: formData
        });
        const result = await response.json();
        formID = result.formID;
        document.getElementById("status").textContent = `Form submitted. Form ID: ${formID}`;
    } catch (error) {
        document.getElementById("status").textContent = "Error submitting form.";
        console.error("Form submission error:", error);
        return;
    }

    const videoFile = form.video.files[0];
    if (videoFile) {
        const videoFormData = new FormData();
        videoFormData.append("video", videoFile);
        videoFormData.append("form_id", formID);

        try {
            const videoResponse = await fetch("/upload-video", {
                method: "POST",
                body: videoFormData
            });
            const videoResult = await videoResponse.json();
            document.getElementById("status").textContent += `\n${videoResult.message}. File URL: ${videoResult.file_url}`;
        } catch (videoError) {
            document.getElementById("status").textContent += "\nError uploading video.";
            console.error("Video upload error:", videoError);
        }
    }
});
