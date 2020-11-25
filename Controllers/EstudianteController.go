package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-template/ApiHelpers"
	"go-api-template/InputFormats"
	"go-api-template/Models"
	"go-api-template/OutputFormats"
	"go-api-template/Repositories"
	"go-api-template/RequestMessages"
	"go-api-template/Utils"
)

/*
	*
	*  FUNCIÓN ListEstudiante
	*
    *
	*
	*
    *
*/

// @Summary Lista de estudiantes
// @Description Lista todos los estudiantes
// @Accept  json
// @Produce  json
// @Success 200 {array} ResponseMessages.ListEstudiantesResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes [get]
func ListEstudiantes(c *gin.Context) {
	// model container
	var container []Models.Estudiante

	// query
	err := Repositories.GetAllEstudiantes(&container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetEstudiantesOutput(container))
}

/*
	*
	*  FUNCIÓN GetOneEstudiante
	*
    *
	*
	*
    *
*/

// @Summary Obtiene un estudiante
// @Description Obtiene un estudiante según su UUID
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a buscar"
// @Success 200 {object} ResponseMessages.GetOneEstudianteResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [get]
func GetOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.GetOneEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.GetOneEstudianteOutput(container))
}

/*
	*
	*  FUNCIÓN AddNewEstudiante
	*
    *
	*
	*
    *
*/

// @Summary Agrega un nuevo estudiante
// @Description Genera un nuevo estudiante con los datos entregados
// @Accept  json
// @Produce  json
// @Param   input_estudiante     body    RequestMessages.AddNewEstudiantePayload     true        "Estudiante a agregar"
// @Success 200 {object} ResponseMessages.AddNewEstudianteResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes [post]
func AddNewEstudiante(c *gin.Context) {
	// input container
	var container RequestMessages.AddNewEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.AddNewEstudianteInput(&container)

	// generate model entity
	model_container := Models.Estudiante{
		Id_rol:                        container.Id_rol,
		Rut_estudiante:                container.Rut_estudiante,
		Nombres_estudiante:            container.Nombres_estudiante,
		Apellidos_estudiante:          container.Apellidos_estudiante,
		Hash_contrasena_estudiante:    container.Hash_contrasena_estudiante,
		Correo_electronico_estudiante: container.Correo_electronico_estudiante,
		Telefono_fijo_estudiante:      container.Telefono_fijo_estudiante,
		Telefono_celular_estudiante:   container.Telefono_celular_estudiante,
	}

	// query
	err := Repositories.AddNewEstudiante(&model_container)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.AddNewEstudianteOutput(model_container))
}

/*
	*
	*  FUNCIÓN PutOneEstudiante
	*
    *
	*
	*
    *
*/

// @Summary Modifica un estudiante
// @Description Modifica un estudiante con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a modificar"
// @Param   input_actualiza_estudiante     body    RequestMessages.PutOneEstudiantePayload     true        "Estudiante a modificar"
// @Success 200 {object} ResponseMessages.PutOneEstudianteResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [put]
func PutOneEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// input container
	var container RequestMessages.PutOneEstudiantePayload

	// input bind
	if err := c.ShouldBind(&container); err != nil {
		ApiHelpers.RespondError(c, 400, "default")
		return
	}

	// format input
	InputFormats.PutOneEstudianteInput(&container)

	// generate model entity
	var model_container Models.Estudiante

	// get query
	err := Repositories.GetOneEstudiante(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// replace data in model entity
	model_container = Models.Estudiante{
		ID:                            model_container.ID,
		Id_rol:                        Utils.CheckUpdatedInt(container.Id_rol, model_container.Id_rol),
		Rut_estudiante:                Utils.CheckUpdatedString(container.Rut_estudiante, model_container.Rut_estudiante),
		Nombres_estudiante:            Utils.CheckUpdatedString(container.Nombres_estudiante, model_container.Nombres_estudiante),
		Apellidos_estudiante:          Utils.CheckUpdatedString(container.Apellidos_estudiante, model_container.Apellidos_estudiante),
		Hash_contrasena_estudiante:    Utils.CheckUpdatedString(container.Hash_contrasena_estudiante, model_container.Hash_contrasena_estudiante),
		Correo_electronico_estudiante: Utils.CheckUpdatedString(container.Correo_electronico_estudiante, model_container.Correo_electronico_estudiante),
		Telefono_fijo_estudiante:      Utils.CheckUpdatedString(container.Telefono_fijo_estudiante, model_container.Telefono_fijo_estudiante),
		Telefono_celular_estudiante:   Utils.CheckUpdatedString(container.Telefono_celular_estudiante, model_container.Telefono_celular_estudiante),
	}

	// put query
	err = Repositories.PutOneEstudiante(&model_container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.PutOneEstudianteOutput(model_container))
}

/*
	*
	*  FUNCIÓN DeleteEstudiante
	*
    *
	*
	*
    *
*/

// @Summary Elimina un estudiante
// @Description Elimina un estudiante con los datos entregados
// @Accept  json
// @Produce  json
// @Param   uuid_estudiante     path    string     true        "UUID del estudiante a eliminar"
// @Success 200 {object} ResponseMessages.DeleteEstudianteResponse "OK"
// @Failure 400 {object} ApiHelpers.ResponseError "Bad request"
// @Router /estudiantes/{uuid_estudiante} [delete]
func DeleteEstudiante(c *gin.Context) {
	// params
	id := c.Params.ByName("id")

	// model container
	var container Models.Estudiante

	// query
	err := Repositories.DeleteEstudiante(&container, id)
	if err != nil {
		ApiHelpers.RespondError(c, 500, "default")
		return
	}

	// output
	ApiHelpers.RespondJSON(c, 200, OutputFormats.DeleteEstudianteOutput(container))
}