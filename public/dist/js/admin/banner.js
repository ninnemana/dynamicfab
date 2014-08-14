define(['alertify'],function(alertify){
	$(document).on('click','.delete',function(e){
		var that = this;
		alertify.confirm('Are you sure you want to delete this banner?',function(e){
			var id = $(that).data('id');
			$.ajax({
				url: '/admin/banners/'+id,
				type: 'DELETE',
				success:function(data, status, xhr){
					$(that).closest('tr').remove();
					alertify.success('Banner removed');
				},
				error: function(xhr, status, err){
					alertify.error(xhr.responseText);
				}
			});
		});
		return false;
	});
});